package userMiddlewares

import (
	"github.com/Doni0414/go-canteen-project/models"
	"github.com/gin-gonic/gin"
)

func IsOrderOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAny, _ := c.Get("user")
		user, _ := userAny.(models.User)

		orderId := c.Query("id")

		if orderId == "" {
			c.Redirect(302, "/")
			c.Abort()
			return
		}

		var order models.Order

		models.DB.Where("id = ?", orderId).First(&order)

		if user.ID != order.UserId {
			c.Redirect(302, "/user/main")
			c.Abort()
			return
		}

		c.Next()
	}
}
