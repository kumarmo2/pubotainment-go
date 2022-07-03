package authentication

import (
	"fmt"
	"net/http"

	authBus "pubwebservice/business/authentication"

	authDtos "pubwebservice/dtos/authentication"

	"github.com/gin-gonic/gin"
)

func RegisterAdmin(ctx *gin.Context) {

	companyName := ctx.Param("companyName")
	if companyName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid companyName"})
		return
	}
	var request authDtos.AdminRegistrationRequest

	if err := ctx.BindJSON(&request); err != nil {
		return
	}

	fmt.Printf("\n\nrequest: %+v\n\n", request)
	if request.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password cannot be empty"})
		return
	}
	err := authBus.RegisterAdmin(companyName, request)
	ctx.JSON(http.StatusOK, err)

}
