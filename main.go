package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dillongc21/RealEstateApp.Backend/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadGoDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
}

func main() {
	loadGoDotEnv()

	router := gin.Default()

	addresses := router.Group("/addresses")
	{
		addresses.GET("/", handler.GetAddresses)
		addresses.GET("/:id", handler.GetAddressById)
		addresses.POST("/", handler.PostAddress)
		addresses.PATCH("/:id/line1", handler.PatchAddressLine1)
	}

	port := os.Getenv("PORT")
	if (port == "") {
		port = "8080"
	}
	runline := fmt.Sprintf("localhost:%s", port)
	router.Run(runline)
}
