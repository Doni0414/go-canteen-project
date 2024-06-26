package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string
	Products *[]Product `gorm:"foreignKey:CategoryId;constraint:OnDelete:CASCADE"`
}
