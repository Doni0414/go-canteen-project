package routes

import (
	"github.com/Doni0414/go-canteen-project/controllers"
	"github.com/Doni0414/go-canteen-project/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.GET("/auth", middlewares.IsAuthorizedAllow(), controllers.AuthPage)
	r.POST("/auth", middlewares.IsAuthorizedAllow(), controllers.Auth)
	r.GET("/auth/logout", controllers.Logout)
}
