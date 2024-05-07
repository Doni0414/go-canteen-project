package userControllers

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/Doni0414/go-canteen-project/dto"
	"github.com/Doni0414/go-canteen-project/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	user, exists := c.Get("user")

	data := gin.H{"title": "Главная страница"}

	q := c.Query("q")
	category := c.Query("category")
	pageStr := c.Query("page")
	pageSize := 20

	if pageStr == "" {
		pageStr = "1"
	}

	page, _ := strconv.Atoi(pageStr)

	var products []models.Product
	var allProducts []models.Product

	if category == "" {
		if q == "" {
			models.DB.Find(&allProducts)
			models.DB.Preload("Category").Limit(pageSize).Offset((page - 1) * pageSize).Find(&products)
		} else {
			models.DB.Where("name LIKE ?", "%"+q+"%").Find(&allProducts)
			models.DB.Where("name LIKE ?", "%"+q+"%").Preload("Category").Limit(pageSize).Offset((page - 1) * pageSize).Find(&products)
		}
	} else {
		categoryId, _ := strconv.Atoi(category)
		if q == "" {
			models.DB.Where("category_id = ?", categoryId).Find(&allProducts)
			models.DB.Where("category_id = ?", categoryId).Preload("Category").Limit(pageSize).Offset((page - 1) * pageSize).Find(&products)
		} else {
			models.DB.Where("category_id = ? AND name LIKE ?", categoryId, "%"+q+"%").Find(&allProducts)
			models.DB.Where("category_id = ? AND name LIKE ?", categoryId, "%"+q+"%").Preload("Category").Limit(pageSize).Offset((page - 1) * pageSize).Find(&products)
		}
	}
	maxPages := int(math.Ceil(float64(len(allProducts)) / float64(pageSize)))

	start := page - 5
	if start < 1 {
		start = 1
	}

	end := page + 5
	if end > maxPages {
		end = maxPages
	}

	result := dto.Result{
		Products:    products,
		CategoryId:  category,
		Query:       q,
		MaxPages:    maxPages,
		FirstPage:   start,
		LastPage:    end,
		CurrentPage: page,
	}

	data["result"] = result

	if exists {
		data["user"] = user
	}

	c.HTML(200, "index.html", data)
}

func AddToCart(c *gin.Context) {
	var cart models.Cart

	if err := c.Bind(&cart); err != nil {
		c.Redirect(302, "/user/main")
		return
	}

	userAny, _ := c.Get("user")

	user, ok := userAny.(models.User)

	if ok {
		cart.UserId = user.ID
		cart.Date = time.Now()

		if cart.Quantity < 1 {
			cart.Quantity = 1
		}

		models.DB.Create(&cart)
	}

	q := c.Query("q")
	category := c.Query("category")
	pageStr := c.Query("page")

	c.Redirect(302, fmt.Sprintf("/user/main?q=%s&category=%s&page=%s", q, category, pageStr))
}
