package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	authService "pubwebservice/services/authentication"
	"pubwebservice/services/events"
	"pubwebservice/services/middlewares"
	"pubwebservice/services/playlists"
	"pubwebservice/services/songs"
)

func startWebService() {
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

	apiGroup.POST("/admin/signin/", authService.SignInAdmin)
	apiGroup.POST("/admin/register/:companyName/", authService.RegisterAdmin)

	adminGroup := apiGroup.Group("/admin", middlewares.AdminAuthMiddleWare)
	{
		adminGroup.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, nil) })
		adminGroup.POST("/playlists/", playlists.AddSong)

		songsGroup := adminGroup.Group("/songs")
		{
			songsGroup.POST("/inventory", songs.AddSongToInventory)
		}
	}
}

func registerUserRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/user/register/:companyName/", authService.RegisterUser)
	apiGroup.POST("/user/signin/", authService.SignInUser)

	userGroup := apiGroup.Group("/user", middlewares.UserAuthMiddleWare)
	{

		userGroup.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, nil)
		})

		userGroup.GET("/events", events.Events)
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
