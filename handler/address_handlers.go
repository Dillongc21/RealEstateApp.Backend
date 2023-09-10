package handler

import (
	"net/http"

	"github.com/Dillongc21/RealEstateApp.Backend/mockdata"
	"github.com/Dillongc21/RealEstateApp.Backend/model"
	"github.com/gin-gonic/gin"
)

func GetAddresses(context *gin.Context) {
	context.JSON(http.StatusOK, mockdata.GetMockAddresses())
}

func PostAddress(context *gin.Context) {
	var newAddress model.Address

	err := context.BindJSON(&newAddress)
	if err != nil { return }

	mockdata.AppendMockAddress(newAddress)
	context.JSON(http.StatusCreated, newAddress)
}


