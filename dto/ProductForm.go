package dto

import "github.com/Doni0414/go-canteen-project/models"

type ProductForm struct {
	models.Product
	ProductID        uint   `form:"id"`
	SearchCategoryId string `form:"search_category_id"`
	Query            string `form:"q"`
	Page             string `form:"page"`
}
