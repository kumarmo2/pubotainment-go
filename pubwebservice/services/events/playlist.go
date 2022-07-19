package events

import (
	"net/http"
	"pubwebservice/business/events"

	"github.com/gin-gonic/gin"
)

func Events(c *gin.Context) {
	events := events.GetEvents()
	c.JSON(http.StatusOK, gin.H{"result": events})
}
