package dto

type OrderStatusForm struct {
	Id      uint   `form:"id"`
	Comment string `form:"comment"`
}
