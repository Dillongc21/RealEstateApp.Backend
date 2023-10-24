package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Dillongc21/RealEstateApp.Backend/constants"
	"github.com/Dillongc21/RealEstateApp.Backend/data"
	"github.com/Dillongc21/RealEstateApp.Backend/data/mock"
	"github.com/Dillongc21/RealEstateApp.Backend/model"
	"github.com/gin-gonic/gin"
)

func GetAddresses(context *gin.Context) {
	context.JSON(http.StatusOK, mock.GetMockAddresses())
}

func GetAddressById(context *gin.Context) {
	id := context.Param("id")
	address := mock.GetMockAddressByID(id)
	if address == (model.Address{}) {
		context.JSON(http.StatusNotFound, gin.H{"message": "Address not found."})
        return
	}
	context.JSON(http.StatusOK, address)
}

func PostAddress(context *gin.Context) {
	var newAddress model.Address

	err := context.BindJSON(&newAddress)
	if err != nil { 
        context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
	mock.AppendMockAddress(newAddress)
	context.JSON(http.StatusCreated, newAddress)
}

func PatchAddressLine1(context *gin.Context) {
	id := context.Param("id")

	newLine1, success := context.GetQuery("new-value")
	if !success {
		context.JSON(http.StatusBadRequest, gin.H{"message": constants.ERROR_MISSING_LINE1})
		return
	}
	updatedAddress, err := mock.UpdateAddressLine1(id, newLine1)

    if err != nil {
        if err.Error() == constants.ERROR_ADDRESS_NOT_FOUND {
            context.JSON(http.StatusNotFound, gin.H{"message": constants.ERROR_ADDRESS_NOT_FOUND})
            return
        }
        if err.Error() == constants.ERROR_MISSING_LINE1 {
            context.JSON(http.StatusBadRequest, gin.H{"message": constants.ERROR_MISSING_LINE1})
            return
        }
        context.JSON(http.StatusBadRequest, gin.H{"message": constants.ERROR_UNKNOWN})
        return
    }
	context.JSON(http.StatusOK, updatedAddress)
}

func DeleteAddress(context *gin.Context) {
    id := context.Param("id")
    success, err := false, errors.New(constants.ERROR_DATASOURCE_NOT_FOUND)

    if *data.GetDataSourceInstance() == "mock" {
        success, err = mock.DeleteAddress(id)
    }
    if err != nil && strings.Contains(err.Error(), "not found"){
        context.JSON(http.StatusNotFound, gin.H{"message": err.Error(), 
                        "datasource": *data.GetDataSourceInstance()})
        return
    }
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    if success {
        context.JSON(http.StatusOK, gin.H{"message": "Address Deleted"})
        return
    }
    context.JSON(http.StatusBadRequest, gin.H{"message": "Unknown error"})
}
    
