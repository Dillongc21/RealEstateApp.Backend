package main

import (
	//"flag"
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

func getPortNumber() string {
	port := os.Getenv("PORT")
	if (port == "") {
		port = "8080"
	}
	return port
}

func main() {
	router := gin.Default()


	addresses := router.Group("/addresses")
	{
		addresses.GET("", handler.GetAddresses)
		addresses.GET("/:id", handler.GetAddressById)
		addresses.POST("", handler.PostAddress)
		addresses.PATCH("/:id/line1", handler.PatchAddressLine1)
        addresses.DELETE("/:id", handler.DeleteAddress)
	}

    //ds := flag.String("datasource", "postgres", "blah blah")
    //flag.Parse()
    //println("Datasource: ", *ds)

	loadGoDotEnv()
	port := getPortNumber()
	runline := fmt.Sprintf("localhost:%s", port)
	router.Run(runline)
}
