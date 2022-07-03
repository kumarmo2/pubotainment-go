package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pubwebservice/business/authentication"
	authService "pubwebservice/services/authentication"
)

func main() {
	fmt.Println("hello world")
	router := gin.Default()
	authentication.X()

	registerRoutes(router)
	http.ListenAndServe(":8000", router)
}

func registerRoutes(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ping")
	})

	router.POST("/register/:companyName/admin/", authService.RegisterAdmin)

}
