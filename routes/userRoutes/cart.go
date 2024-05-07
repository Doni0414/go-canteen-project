package userRoutes

import (
	"github.com/Doni0414/go-canteen-project/controllers/userControllers"
	"github.com/Doni0414/go-canteen-project/middlewares"
	"github.com/Doni0414/go-canteen-project/middlewares/userMiddlewares"
	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine) {
	r.POST("/user/cart/addToCart", middlewares.IsAuthorized(), userMiddlewares.IsUser(), userControllers.AddToCart)
	r.GET("/user/cart", middlewares.IsAuthorized(), userMiddlewares.IsUser(), userControllers.CartPage)
	r.POST("/user/cart/removeFromCart", middlewares.IsAuthorized(), userMiddlewares.IsUser(), userControllers.RemoveFromCart)
	r.POST("/user/cart/changeQuantity", middlewares.IsAuthorized(), userMiddlewares.IsUser(), userControllers.ChangeQuantity)
}
