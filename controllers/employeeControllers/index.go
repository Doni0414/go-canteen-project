package employeeControllers

import (
	"fmt"
	"math"
	"path/filepath"
	"strconv"

	"github.com/Doni0414/go-canteen-project/dto"
	"github.com/Doni0414/go-canteen-project/models"
	"github.com/Doni0414/go-canteen-project/utils"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	q := c.Query("q")
	category := c.Query("category")
	pageStr := c.Query("page")
	fmt.Println(q, category, pageStr)
	data := retrieveMainPageData(c, q, category, pageStr)
	c.HTML(200, "employeeIndex.html", data)
}

var IMAGE_UPLOAD_PATH = "C:/Users/сулпак/go/go-canteen-project/static/images/"
var RELATIVE_IMAGE_PATH = "/static/images/"

func CreateProduct(c *gin.Context) {
	var productForm dto.ProductForm
	c.Bind(&productForm)
	data := gin.H{}

	if err := productForm.Validate(); err != nil {
		data["product_creation_error"] = err.Error()
		c.JSON(400, data)
		return
	}

	image, err := c.FormFile("image")

	if err != nil {
		data["product_creation_error"] = "Выберите картинку!"
		c.JSON(400, data)
		return
	}

	baseName := filepath.Base(image.Filename)
	generatedName := utils.GenerateUniqueName(baseName)
	productForm.Image = RELATIVE_IMAGE_PATH + generatedName
	fmt.Println(productForm)
	if err := c.SaveUploadedFile(image, IMAGE_UPLOAD_PATH+generatedName); err != nil {
		data["product_creation_error"] = err.Error()
		c.JSON(400, data)
		return
	}

	models.DB.Create(&productForm.Product)
}

func GetProduct(c *gin.Context) {
	data := gin.H{}

	id := c.Query("id")

	fmt.Println("id = ", id)

	var product models.Product
	models.DB.Where("id = ?", id).Preload("Category").First(&product)

	data["product"] = product

	c.JSON(200, data)
}

func EditProduct(c *gin.Context) {
	var productForm dto.ProductForm
	c.Bind(&productForm)
	data := gin.H{}

	if err := productForm.Validate(); err != nil {
		data["product_edit_error"] = err.Error()
		c.JSON(400, data)
		return
	}

	image, err := c.FormFile("image")

	if err != nil {
		var product models.Product
		models.DB.Where("id = ?", productForm.ProductID).First(&product)
		productForm.Image = product.Image
	} else {
		baseName := filepath.Base(image.Filename)
		generatedName := utils.GenerateUniqueName(baseName)
		productForm.Image = RELATIVE_IMAGE_PATH + generatedName
		fmt.Println(productForm)
		if err := c.SaveUploadedFile(image, IMAGE_UPLOAD_PATH+generatedName); err != nil {
			data["product_edit_error"] = err.Error()
			c.JSON(400, data)
			return
		}
	}

	productForm.ID = productForm.ProductID
	models.DB.Save(&productForm.Product)
}

func DeleteProduct(c *gin.Context) {
	var productForm dto.ProductForm
	c.Bind(&productForm)

	fmt.Println(productForm.ProductID)

	models.DB.Delete(&models.Product{}, "id = ?", productForm.ProductID)

	c.Redirect(302, fmt.Sprintf("/employee/main?q=%s&category=%s&page=%s", productForm.Query, productForm.SearchCategoryId, productForm.Page))
}

func retrieveMainPageData(c *gin.Context, q, category, pageStr string) gin.H {
	data := gin.H{"title": "Главная страница"}

	user, exists := c.Get("user")
	if exists {
		data["user"] = user
	}

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

	return data
}
