package employeeRoutes

import (
	"github.com/Doni0414/go-canteen-project/controllers/employeeControllers"
	"github.com/Doni0414/go-canteen-project/middlewares"
	"github.com/Doni0414/go-canteen-project/middlewares/employeeMiddlewares"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.GET("/employee/orders", middlewares.IsAuthorized(), employeeMiddlewares.IsEmployee(), employeeControllers.OrdersPage)
	r.GET("/employee/orders/order", middlewares.IsAuthorized(), employeeMiddlewares.IsEmployee(), employeeControllers.OrderPage)
	r.POST("/employee/orders/order/next", middlewares.IsAuthorized(), employeeMiddlewares.IsEmployee(), employeeControllers.NextOrderStatus)
	r.POST("/employee/orders/order/reject", middlewares.IsAuthorized(), employeeMiddlewares.IsEmployee(), employeeControllers.RejectOrderStatus)

}
