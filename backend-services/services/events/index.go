package events

import (
	"fmt"
	"log"
	"time"

	"encoding/json"
	qu "pubwebservice/business/queue_utils"
	qm "pubwebservice/commonLibs/queue_manager"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func payloadHandler(conn *websocket.Conn, c *gin.Context, payload interface{}) {
	fmt.Printf("payload received: %v", payload)
	bytes, _ := json.Marshal(payload)
	conn.WriteMessage(1, bytes)
}

func generateHandler(conn *websocket.Conn, c *gin.Context) func(interface{}) {
	return func(payload interface{}) {
		payloadHandler(conn, c, payload)
	}
}

func WsHandler(c *gin.Context) {
	/*
	   - get the deviceId & companyId from the request.
	   - for a company, get the queueManager.
	   - all the devices that belong to the same company, which got connected to the same server,
	       (they might get connected to different servers), will get the same instance of QueueManager,
	       So that the notifications could be sent to all of them.
	   - add as a subscriber to the queue.
	   - wait for receiving the message on the queue.


	*/

	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	deviceId := c.Keys["deviceId"]
	companyId := c.Keys["companyId"]
	log.Printf("deviceId: %v\n, companyId: %v\n", deviceId, companyId)
	qManager := qu.GetQueueManagerFactory().GetQueueManager(fmt.Sprint(companyId))

	handler := generateHandler(conn, c)
	sub := qm.NewSubscriber(handler)
	qManager.AddSubscriber(sub)
	defer qManager.RemoveSubscriber(sub)

	go func() {
		for {
			qManager.BroadCast("msg...")
			time.Sleep(time.Second * 3)
		}

	}()

	conn.ReadMessage() // NOTE: this will block, until client sends any message. no client should be sending
	// message to the server.

	// for {
	// if err != nil {
	// break
	// }

	// msgString := string(msg)

	// fmt.Printf("message received: %v", msgString)
	// fmt.Printf("messageType: %v", messageType)
	// }

}
