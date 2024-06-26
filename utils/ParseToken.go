package utils

import (
	"github.com/Doni0414/go-canteen-project/models"
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("asdfasffbgdfbfsvsdvcxzvsfvsdvadwegrthtyjuyjdgbxfvbadqwdqwdsvbgnjytujuykimgfbdsdsdf1233543"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
