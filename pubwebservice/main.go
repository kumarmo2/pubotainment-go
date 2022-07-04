package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	authService "pubwebservice/services/authentication"
)

func main() {
	fmt.Println("hello world")
	router := gin.Default()

	registerRoutes(router)
	http.ListenAndServe(":8000", router)
}

func registerRoutes(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ping")
	})

	var registerRoute = router.Group("/register")
	{
		registerRoute.POST("/:companyName/admin/", authService.RegisterAdmin)
		registerRoute.POST("/:companyName/user/", authService.RegisterUser)
	}
}
