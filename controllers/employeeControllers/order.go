package employeeControllers

import (
	"fmt"
	"strconv"

	"github.com/Doni0414/go-canteen-project/dto"
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

	models.DB.Preload("OrderDetails").Preload("OrderDetails.Product").Order("created_at desc").Find(&orders)

	data["orders"] = orders

	c.HTML(200, "employeeOrders.html", data)
}

func OrderPage(c *gin.Context) {
	data := gin.H{}

	userAny, _ := c.Get("user")
	user, _ := userAny.(models.User)
	data["user"] = user

	orderId := c.Query("id")

	if orderId == "" {
		c.Redirect(302, "/employee/orders")
		return
	}

	var order models.Order
	models.DB.Where("id = ?", orderId).Preload("OrderHistories", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC")
	}).Preload("OrderDetails").Preload("OrderDetails.Product").First(&order)

	data["order"] = order

	c.HTML(200, "employeeOrder.html", data)
}

func NextOrderStatus(c *gin.Context) {
	var form dto.OrderStatusForm

	if err := c.Bind(&form); err != nil {
		fmt.Println(err)
		c.Redirect(302, "/employee/orders")
		return
	}

	var order models.Order
	models.DB.Where("id = ?", form.Id).First(&order)

	if order.ID == 0 {
		fmt.Println("no such order")
		c.Redirect(302, "/employee/orders")
		return
	}

	if order.IsFinished() {
		c.Redirect(302, "/employee/orders/order?id="+strconv.Itoa(int(order.ID)))
		return
	}

	order.ChangeStatus()
	models.DB.Save(order)

	orderHistory := models.OrderHistory{
		OrderId: order.ID,
		Status:  order.Status,
		Comment: form.Comment,
	}

	models.DB.Create(&orderHistory)

	c.Redirect(302, "/employee/orders/order?id="+strconv.Itoa(int(order.ID)))
}

func RejectOrderStatus(c *gin.Context) {
	var form dto.OrderStatusForm

	if err := c.Bind(&form); err != nil {
		fmt.Println(err)
		c.Redirect(302, "/employee/orders")
		return
	}

	var order models.Order
	models.DB.Where("id = ?", form.Id).First(&order)

	if order.ID == 0 {
		fmt.Println("no such order")
		c.Redirect(302, "/employee/orders")
		return
	}

	if order.IsFinished() {
		c.Redirect(302, "/employee/orders/order?id="+strconv.Itoa(int(order.ID)))
		return
	}

	order.Reject()
	models.DB.Save(order)

	orderHistory := models.OrderHistory{
		OrderId: order.ID,
		Status:  order.Status,
		Comment: form.Comment,
	}

	models.DB.Create(&orderHistory)

	c.Redirect(302, "/employee/orders/order?id="+strconv.Itoa(int(order.ID)))
}
