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

func helloMiddleWare(c *gin.Context) {
	fmt.Println("hello from middleware")
	c.JSON(http.StatusOK, "returning from middleware")
	c.Abort()
}

func registerAdminRoutes(apiGroup *gin.RouterGroup) {

	adminGroup := apiGroup.Group("/admin")
	{
		adminGroup.POST("/signin/", authService.SignInAdmin)
		adminGroup.POST("/register/:companyName/", authService.RegisterAdmin)
	}
}

func registerUserRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/user/register/:companyName/", authService.RegisterUser)
	apiGroup.POST("/user/signin/", authService.SignInUser)

	userGroup := apiGroup.Group("/user")
	{
		userGroup.POST("/ping", func(ctx *gin.Context) {})
	}
}

func registerRoutes(router *gin.Engine) {

	var apiGroup = router.Group("/api")
	{
		apiGroup.GET("/", helloMiddleWare, func(ctx *gin.Context) {
			fmt.Println("hello from root")
			ctx.JSON(http.StatusOK, "ping")
		})

		apiGroup.GET("/props", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"name": "kumarmo2"})
		})

		registerAdminRoutes(apiGroup)
		registerUserRoutes(apiGroup)

	}

}
