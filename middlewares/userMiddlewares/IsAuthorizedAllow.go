package userMiddlewares

import (
	"github.com/Doni0414/go-canteen-project/models"
	"github.com/Doni0414/go-canteen-project/utils"
	"github.com/gin-gonic/gin"
)

func IsAuthorizedAllow() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")

		if err != nil {
			c.Next()
			return
		}

		claims, err := utils.ParseToken(token)

		if err != nil {
			c.Next()
			return
		}

		var user models.User

		models.DB.Where("email = ?", claims.Subject).First(&user)
		if user.ID == 0 {
			c.Next()
			return
		}

		if user.Role != "user" {
			c.Redirect(302, "/auth")
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Set("user", user)
		c.Next()
	}
}
