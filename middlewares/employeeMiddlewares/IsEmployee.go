package employeeMiddlewares

import (
	"github.com/Doni0414/go-canteen-project/models"
	"github.com/gin-gonic/gin"
)

func IsEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAny, _ := c.Get("user")
		user, _ := userAny.(models.User)

		if user.Role != "employee" {
			c.Redirect(302, "/user/main")
			c.Abort()
			return
		}

		c.Next()
	}
}
