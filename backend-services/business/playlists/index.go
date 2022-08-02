package playlists

import (
	"errors"
	// queueUtils "pubwebservice/business/queue_utils"
	"pubwebservice/dtos/songs"
)

func AddToPlaylist(request *songs.AddToPlaylistRequest, companyId int64) error {
	if request == nil {
		panic("Invalid request")
	}
	if request.Name == "" {
		return errors.New("Song name cannot be empty")
	}
	// queueManager := queueUtils.GetQueueManagerFactory().GetQueueManager()
	// queueManager.BroadCast(request)

	return nil
}
