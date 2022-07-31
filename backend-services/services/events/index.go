package events

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WsHandler(c *gin.Context) {

	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		msgString := string(msg)

		fmt.Printf("message received: %v", msgString)
		fmt.Printf("messageType: %v", messageType)
	}

}
