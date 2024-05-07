package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Role string
	jwt.StandardClaims
}
