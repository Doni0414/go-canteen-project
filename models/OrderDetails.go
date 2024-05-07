package models

import "gorm.io/gorm"

type OrderDetails struct {
	gorm.Model
	OrderId   uint
	ProductId uint
	Quantity  uint
	Product   Product `gorm:"foreignkey:ProductId"`
}

func (od OrderDetails) TotalPrice() float32 {
	return od.Product.Price * float32(od.Quantity)
}
