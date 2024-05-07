package dto

import "github.com/Doni0414/go-canteen-project/models"

type Result struct {
	Products    []models.Product
	CategoryId  string
	Query       string
	MaxPages    int
	FirstPage   int
	LastPage    int
	CurrentPage int
}

func (r Result) PreviousPage() int {
	return r.CurrentPage - 1
}

func (r Result) NextPage() int {
	return r.CurrentPage + 1
}

func (r Result) IsLastPage() bool {
	return r.CurrentPage == r.MaxPages
}

func (r Result) IsFirstPage() bool {
	return r.CurrentPage == 1
}

func (r Result) CategoryName() string {
	if r.CategoryId == "" {
		return "Все"
	}
	var category models.Category
	models.DB.Where("id = ?", r.CategoryId).First(&category)
	return category.Name
}
