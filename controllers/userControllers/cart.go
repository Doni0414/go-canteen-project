package userControllers

import (
	"fmt"

	"github.com/Doni0414/go-canteen-project/dto"
	"github.com/Doni0414/go-canteen-project/models"
	"github.com/gin-gonic/gin"
)

func CartPage(c *gin.Context) {
	data := gin.H{}

	userAny, _ := c.Get("user")
	user, _ := userAny.(models.User)
	data["user"] = user

	var cart []models.Cart

	models.DB.Where("user_id = ?", user.ID).Preload("Product").Find(&cart)

	result := dto.CartResult{
		Cart: cart,
	}

	data["CartResult"] = result

	c.HTML(200, "cart.html", data)
}

func RemoveFromCart(c *gin.Context) {
	var cart dto.CartRemove

	c.Bind(&cart)

	models.DB.Delete(&models.Cart{}, "id = ?", cart.Id)

	c.Redirect(302, "/user/cart")
}

func ChangeQuantity(c *gin.Context) {
	var cartQuantity dto.CartQuantity

	if err := c.ShouldBindJSON(&cartQuantity); err != nil {
		fmt.Println(err)
	}

	var cart models.Cart

	models.DB.Where("id = ?", cartQuantity.Id).Preload("Product").First(&cart)

	cart.Quantity = cartQuantity.Quantity

	models.DB.Save(&cart)

	data := gin.H{}
	data["product-price"] = cart.Product.Price
	data["product-total-price"] = cart.TotalPrice()

	userAny, _ := c.Get("user")
	user, _ := userAny.(models.User)
	data["user"] = user

	var allCarts []models.Cart

	models.DB.Where("user_id = ?", user.ID).Preload("Product").Find(&allCarts)

	result := dto.CartResult{
		Cart: allCarts,
	}

	data["total-price"] = result.TotalPrice()

	c.JSON(200, data)
}

func MakeOrder(c *gin.Context) {
	userAny, _ := c.Get("user")

	user, _ := userAny.(models.User)

	var carts []models.Cart

	models.DB.Where("user_id = ?", user.ID).Find(&carts)

	if len(carts) > 0 {
		order := models.Order{
			UserId: user.ID,
			Status: "В очереди",
		}

		models.DB.Create(&order)

		orderHistory := models.OrderHistory{
			OrderId: order.ID,
			Status:  "В очереди",
		}

		models.DB.Create(&orderHistory)

		var orderDetails []models.OrderDetails

		for _, cart := range carts {
			orderDetail := models.OrderDetails{
				OrderId:   order.ID,
				ProductId: cart.ProductId,
				Quantity:  cart.Quantity,
			}
			orderDetails = append(orderDetails, orderDetail)
		}
		models.DB.Create(&orderDetails)

		models.DB.Delete(&carts)

	}
	c.Redirect(302, "/user/orders")
}
