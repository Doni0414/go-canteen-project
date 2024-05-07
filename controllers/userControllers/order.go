package userControllers

import (
	"github.com/Doni0414/go-canteen-project/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrdersPage(c *gin.Context) {
	data := gin.H{}

	userAny, _ := c.Get("user")
	user, _ := userAny.(models.User)
	data["user"] = user

	var orders []models.Order

	models.DB.Where("user_id = ?", user.ID).Preload("OrderDetails").Preload("OrderDetails.Product").Order("created_at desc").Find(&orders)

	data["orders"] = orders

	c.HTML(200, "orders.html", data)
}

func OrderPage(c *gin.Context) {
	data := gin.H{}

	userAny, _ := c.Get("user")
	user, _ := userAny.(models.User)
	data["user"] = user

	orderId := c.Query("id")

	var order models.Order

	models.DB.Where("id = ?", orderId).Preload("OrderHistories", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC")
	}).Preload("OrderDetails").Preload("OrderDetails.Product").First(&order)

	data["order"] = order

	c.HTML(200, "order.html", data)
}
