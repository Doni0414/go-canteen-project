package employeeRoutes

import (
	"github.com/Doni0414/go-canteen-project/controllers/employeeControllers"
	"github.com/Doni0414/go-canteen-project/middlewares"
	"github.com/Doni0414/go-canteen-project/middlewares/employeeMiddlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	OrderRoutes(r)
	InitIndex(r)
}

func InitIndex(r *gin.Engine) {
	r.GET("/employee/main", middlewares.IsAuthorized(), employeeMiddlewares.IsEmployee(), employeeControllers.Index)
	r.POST("/employee/products/create", middlewares.IsAuthorized(), employeeMiddlewares.IsEmployee(), employeeControllers.CreateProduct)
	r.GET("/employee/products", middlewares.IsAuthorized(), employeeMiddlewares.IsEmployee(), employeeControllers.GetProduct)
	r.POST("/employee/products/edit", middlewares.IsAuthorized(), employeeMiddlewares.IsEmployee(), employeeControllers.EditProduct)
	r.POST("/employee/products/delete", middlewares.IsAuthorized(), employeeMiddlewares.IsEmployee(), employeeControllers.DeleteProduct)
}
