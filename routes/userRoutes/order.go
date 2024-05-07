package userRoutes

import (
	"github.com/Doni0414/go-canteen-project/controllers/userControllers"
	"github.com/Doni0414/go-canteen-project/middlewares"
	"github.com/Doni0414/go-canteen-project/middlewares/userMiddlewares"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.GET("/user/orders", middlewares.IsAuthorized(), userMiddlewares.IsUser(), userControllers.OrdersPage)
	r.POST("/user/order/makeOrder", middlewares.IsAuthorized(), userMiddlewares.IsUser(), userControllers.MakeOrder)
	r.GET("/user/orders/order", middlewares.IsAuthorized(), userMiddlewares.IsUser(), userMiddlewares.IsOrderOwner(), userControllers.OrderPage)
}
