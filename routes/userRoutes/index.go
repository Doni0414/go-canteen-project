package userRoutes

import (
	"github.com/Doni0414/go-canteen-project/controllers/userControllers"
	"github.com/Doni0414/go-canteen-project/middlewares/userMiddlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	InitIndex(r)
	CartRoutes(r)
	OrderRoutes(r)
}

func InitIndex(r *gin.Engine) {
	r.GET("/user/main", userMiddlewares.IsAuthorizedAllow(), userControllers.Index)
}
