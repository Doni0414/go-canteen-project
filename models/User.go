package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `form:"firstname"`
	LastName  string `form:"lastname"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	Role      string
}

type LoginUser struct {
	Email    string `form:"login-email"`
	Password string `form:"login-password"`
}

func (lu *LoginUser) Validate() (err error) {
	if len(strings.TrimSpace(lu.Email)) == 0 || !IsValidEmail(lu.Email) {
		err = errors.New("Неправильная почта!")
		return
	}
	if len(strings.TrimSpace(lu.Password)) == 0 {
		err = errors.New("Пароль не может быть пустым!")
		return
	}
	err = nil
	return
}

type SignupUser struct {
	User
	ConfirmationPassword string `form:"confirmation-password"`
}

func (su *SignupUser) GetUser() User {
	return su.User
}

func (su *SignupUser) Validate() (err error) {
	if len(strings.TrimSpace(su.FirstName)) == 0 {
		err = errors.New("Имя не может быть пустым!")
		return
	}
	if len(strings.TrimSpace(su.LastName)) == 0 {
		err = errors.New("Фамилия не может быть пустым!")
		return
	}

	if len(strings.TrimSpace(su.Email)) == 0 || !IsValidEmail(su.Email) {
		err = errors.New("Неправильная почта!")
		return
	}
	if len(strings.TrimSpace(su.Password)) == 0 {
		err = errors.New("Пароль не может быть пустым!")
		return
	}
	if su.ConfirmationPassword != su.Password {
		err = errors.New("Пароли не совпадают!")
		return
	}

	err = nil
	return
}
