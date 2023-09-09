package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// This handler will match /user/john but will not match /user/ or /user
func getUserByNameHandler(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, "Hello %s", name)
}

func main() {
	router := gin.Default()

	router.GET("/user/:name", getUserByNameHandler)

	router.Run()
}
