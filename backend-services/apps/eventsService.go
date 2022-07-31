package main

import (
	"fmt"
	"net/http"

	"pubwebservice/services/events"

	"github.com/gin-gonic/gin"
)

func startEventsService() {
	fmt.Println("======== events service ===========")
	router := gin.Default()
	router.GET("/events/", events.WsHandler)

	http.ListenAndServe(":8001", router)

}
