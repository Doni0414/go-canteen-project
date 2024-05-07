package models

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// var statuses = []string{"В очереди", "Отклонен", "Приготовление", "В доставке"}

type Order struct {
	gorm.Model
	UserId         uint
	OrderDetails   []OrderDetails `gorm:"foreignkey:OrderId"`
	OrderHistories []OrderHistory `gorm:"foreignkey:OrderId"`
	Status         string
}

func (o Order) GetStatusCSSClass() string {
	if o.Status == "В очереди" {
		return "status-pending"
	} else if o.Status == "Отклонен" {
		return "status-rejected"
	} else if o.Status == "Приготовление" {
		return "status-preparing"
	} else if o.Status == "В доставке" {
		return "status-delivery"
	}
	return "status-delivered"
}

func (o Order) GetProducts() string {
	var productNames []string
	for _, orderDetail := range o.OrderDetails {
		productNames = append(productNames, orderDetail.Product.Name+" x"+strconv.Itoa(int(orderDetail.Quantity)))
	}
	return strings.Join(productNames, ", ")
}

func (o Order) GetDate() string {
	return fmt.Sprintf("%s.%s.%s %s:%s", format(o.CreatedAt.Day()), format(int(o.CreatedAt.Month())), format(o.CreatedAt.Year()), format(o.CreatedAt.Hour()), format(o.CreatedAt.Minute()))
}

func (o Order) TotalPrice() float32 {
	var sum float32 = 0
	for _, od := range o.OrderDetails {
		sum += od.TotalPrice()
	}
	return sum
}

func (o *Order) ChangeStatus() {
	if o.Status == "В очереди" {
		o.Status = "Приготовление"
	} else if o.Status == "Приготовление" {
		o.Status = "В доставке"
	} else if o.Status == "В доставке" {
		o.Status = "Доставлен"
	}
}

func (o *Order) Reject() {
	o.Status = "Отклонен"
}

func (o Order) IsFinished() bool {
	return o.Status == "Доставлен" || o.Status == "Отклонен"
}
