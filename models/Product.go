package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name       string  `form:"name"`
	Price      float32 `form:"price"`
	NumOfStock int
	CategoryId int      `form:"category_id"`
	Category   Category `gorm:"foreignkey:CategoryId"`
	Image      string
}

func (p *Product) Validate() (err error) {
	if len(strings.TrimSpace(p.Name)) == 0 {
		err = errors.New("Имя не может быть пустым!")
		return
	}

	if p.Price <= 0 {
		err = errors.New("Цена не может быть меньше нуля!")
		return
	}

	err = nil
	return
}
