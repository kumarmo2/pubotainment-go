package events

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"encoding/json"
	qu "pubwebservice/business/queue_utils"
	qm "pubwebservice/commonLibs/queue_manager"
	sdDataAccess "pubwebservice/dataAccess/serviceDiscovery"
	eventsDto "pubwebservice/dtos/events"
	sdModel "pubwebservice/models/serviceDiscovery"
	"pubwebservice/services/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func ForwardEventToDevices(c *gin.Context) {
	var request eventsDto.ForwardEventRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if request.Payload == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload cannot be null"})
		c.Abort()
		return
	}

	if request.CompanyId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid companyId"})
		c.Abort()
		return
	}
	qm := qu.GetQueueManagerFactory().GetQueueManager(fmt.Sprintf("%v", request.CompanyId))
	qm.BroadCast(gin.H{"event": request.Payload})
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
	// deviceId := c.Keys["deviceId"]
	companyId := c.Keys["companyId"].(int64)
	qManager := qu.GetQueueManagerFactory().GetQueueManager(fmt.Sprint(companyId))
	connectionId := uuid.New().String()

	handler := generateHandler(conn, c)
	sub := qm.NewSubscriber(handler)
	qManager.AddSubscriber(sub)
	defer qManager.RemoveSubscriber(sub)

	quitChan := make(chan bool)

	go func() {
		for {
			select {
			case <-quitChan:
				{
					return
				}
			default:
				log.Println("ping service discovery")
				// TODO: update the lastpinged of the connection in the cassandra.
				time.Sleep(time.Second * 3)
				connectionMap := sdModel.ConnectionServerMap{
					ServerId:     utils.GetServerId(),
					ConnectionId: connectionId,
					LastPinged:   time.Now(),
					CompanyId:    companyId,
				}
				sdDataAccess.InsertConnectionMap(&connectionMap)
			}
		}

	}()

	conn.ReadMessage() // NOTE: this will block, until client sends any message. no client should be sending
	quitChan <- true
}
