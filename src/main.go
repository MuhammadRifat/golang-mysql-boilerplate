package main

import (
	"log"
	"os"
	"url-shortner/src/config"
	"url-shortner/src/routes"
	"url-shortner/src/seed"
	"url-shortner/src/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	var dbErr error
	config.DB, dbErr = config.ConnectDB()
	if dbErr != nil {
		log.Fatalf("Could not connect to the database: %v", dbErr)
	}
	defer func() {
		dbInstance, _ := config.DB.DB()
		_ = dbInstance.Close()
	}()

	// global middleware
	router := gin.Default()
	router.Use(config.CORS())
	router.Use(util.GlobalErrorHandler())

	// all routes
	routes.Routes(router)

	// db auto migration
	seed.MigrateDB()

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
