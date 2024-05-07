package models

import (
	"fmt"

	"gorm.io/gorm"
)

type OrderHistory struct {
	gorm.Model
	OrderId uint
	Status  string
	Comment string
}

func (oh OrderHistory) HasEmptyComment() bool {
	return oh.Comment == ""
}

func (oh OrderHistory) GetDate() string {
	return fmt.Sprintf("%s.%s.%s в %s:%s", format(oh.CreatedAt.Day()), format(int(oh.CreatedAt.Month())), format(oh.CreatedAt.Year()), format(oh.CreatedAt.Hour()), format(oh.CreatedAt.Minute()))
}
