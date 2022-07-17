package playlists

import (
	"log"
	"pubwebservice/business/playlists"
	"pubwebservice/dtos/songs"

	"github.com/gin-gonic/gin"
)

func AddSong(c *gin.Context) {
	var request songs.AddToPlaylistRequest

	if err := c.BindJSON(&request); err != nil {
		log.Println("err: ", err.Error())
		return
	}
	// companyId := c.Keys["companyId"].(int64)
	playlists.AddToPlaylist(&request, 1)
}
