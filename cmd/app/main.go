package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vietanh-it/go-blog-api/config"
	"github.com/vietanh-it/go-blog-api/pkg/category/controllers"
	"github.com/vietanh-it/go-blog-api/pkg/category/models"
	"log"
	"os"
)

func main() {
	// Env loading
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Databases
	db, dbError := config.ConnectDB()
	if dbError != nil {
		log.Fatal("Error connecting DB")
	}
	var allModels = []interface{}{&models.Category{}}

	// Migration
	migrateErr := db.AutoMigrate(allModels...)
	if migrateErr != nil {
		log.Fatal("Error migrating DB")
	}

	// Routers
	router := gin.Default()
	var category = controllers.CategoryController{}
	category.Handler(router)

	serverErr := router.Run("0.0.0.0:" + os.Getenv("PORT"))
	if serverErr != nil {
		log.Fatal("Server error")
	}
}
