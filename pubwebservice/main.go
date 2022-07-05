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

		signInGroup := apiGroup.Group("/signin")
		{
			signInGroup.POST("/admin/", authService.SignInAdmin)
		}

		var registerRoute = apiGroup.Group("/register")
		{
			registerRoute.POST("/:companyName/admin/", authService.RegisterAdmin)
			registerRoute.POST("/:companyName/user/", authService.RegisterUser)
		}
	}

}
