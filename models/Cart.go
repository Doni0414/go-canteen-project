package models

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint
	ProductId uint `form:"product-id"`
	Quantity  uint `form:"quantity"`
	Date      time.Time
	Product   Product `gorm:"foreignkey:ProductId"`
}

func (c Cart) TotalPrice() float32 {
	return c.Product.Price * float32(c.Quantity)
}

func (c Cart) GetDate() string {
	return fmt.Sprintf("%s.%s.%s %s:%s", format(c.Date.Day()), format(int(c.Date.Month())), format(c.Date.Year()), format(c.Date.Hour()), format(c.Date.Minute()))
}

func format(num int) string {
	str := strconv.Itoa(num)
	if num < 10 {
		return "0" + str
	}
	return str
}
