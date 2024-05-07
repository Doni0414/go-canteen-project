package main

import (
	"log"
	"os"

	"github.com/Doni0414/go-canteen-project/models"
	"github.com/Doni0414/go-canteen-project/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	routes.InitRoutes(r)
	r.LoadHTMLGlob("templates/*/*.html")
	r.Static("/static", "static")

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file %v", err)
	}

	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	models.InitDB(config)

	r.Run(":8080")
}
