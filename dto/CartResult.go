package dto

import "github.com/Doni0414/go-canteen-project/models"

type CartResult struct {
	Cart []models.Cart
}

func (c CartResult) TotalPrice() float32 {
	var sum float32 = 0
	for _, cart := range c.Cart {
		sum += cart.TotalPrice()
	}
	return sum
}
