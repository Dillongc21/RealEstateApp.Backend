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

func GetAddressById(context *gin.Context) {
	id := context.Param("id")
	address := mockdata.GetMockAddressByID(id)
	if address == (model.Address{}) {
		context.JSON(http.StatusNotFound, gin.H{"message": "Address not found."})
	}
	context.JSON(http.StatusOK, address)
}

func PostAddress(context *gin.Context) {
	var newAddress model.Address

	err := context.BindJSON(&newAddress)
	if err != nil { return }

	mockdata.AppendMockAddress(newAddress)
	context.JSON(http.StatusCreated, newAddress)
}

func PatchAddressLine1(context *gin.Context) {
	id := context.Param("id")

	newLine1, success := context.GetQuery("new-value")
	if !success {
		context.JSON(http.StatusBadRequest, gin.H{"message": "missing line1 value query parameter"})
		return
	}

	index, err := mockdata.GetAddressIndex(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "id parameter is not a valid number"})
		return
	}
	if index == -1 {
		context.JSON(http.StatusNotFound, gin.H{"message": "Address not found"})
		return
	}

	updatedAddress, _ := mockdata.UpdateAddressLine1(id, newLine1)
	context.JSON(http.StatusOK, updatedAddress)
}
