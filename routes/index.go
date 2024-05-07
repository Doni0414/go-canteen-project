package routes

import (
	"text/template"

	"github.com/Doni0414/go-canteen-project/middlewares"
	"github.com/Doni0414/go-canteen-project/routes/employeeRoutes"
	"github.com/Doni0414/go-canteen-project/routes/userRoutes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	AuthRoutes(r)
	userRoutes.InitRoutes(r)
	employeeRoutes.InitRoutes(r)
	r.Use(middlewares.CorsMiddleware())
	r.SetFuncMap(template.FuncMap{
		"sequence": sequence,
	})
}

func sequence(start, end int) []int {
	seq := make([]int, end-start+1)
	for i := range seq {
		seq[i] = start + i
	}
	return seq
}
