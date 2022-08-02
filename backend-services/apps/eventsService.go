package main

import (
	"fmt"
	"log"

	"net/http"

	"pubwebservice/commonLibs/misc"
	"pubwebservice/dataAccess/serviceDiscovery"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"pubwebservice/services/events"
	"pubwebservice/services/middlewares"
)

var ServerId uuid.UUID = uuid.New()

func startEventsService() {
	fmt.Println("======== events service ===========")
	addr, err := misc.GetLocalIp()
	if err != nil {
		log.Println("err:", err.Error())
	}
	log.Printf("Local address: %v\n", addr)
	servicediscovery.InsertServcerInstance(ServerId.String(), []string{addr})
	// ips := servicediscovery.GetIps("id1")

	// fmt.Printf("ips: %v\n", ips)

	router := gin.Default()
	// TODO: even admins can listen to events. So either the userAuth or adminAuth should be passed.
	router.GET("/events/", middlewares.UserAuthMiddleWare, events.WsHandler)

	http.ListenAndServe(":8001", router)

}
