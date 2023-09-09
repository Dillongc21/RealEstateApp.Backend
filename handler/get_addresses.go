package handler

import (
	"net/http"
	"github.com/Dillongc21/RealEstateApp.Backend/mockdata"
	"github.com/gin-gonic/gin"
)

func GetAddresses(context *gin.Context) {
	context.JSON(http.StatusOK, mockdata.GetMockAddresses())
}
