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

	router.GET("/addresses", handler.GetAddresses)
	router.POST("/addresses", handler.PostAddress)

	port := os.Getenv("PORT")
	if (port == "") {
		port = "8080"
	}
	runline := fmt.Sprintf("localhost:%s", port)
	router.Run(runline)
}
