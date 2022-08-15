package songs

import (
	"errors"
	"log"
	idGen "pubwebservice/commonLibs/IdGenerator"
	songsDA "pubwebservice/dataAccess/songs"
	"pubwebservice/dtos/songs"
	songModels "pubwebservice/models/songs"
	"time"
)

func AddSongToInventory(request *songs.AddSongToInventory) (int64, error) {
	if request == nil {
		return 0, errors.New("request cannot be null")
	}
	songName := request.Name
	companyId := request.CompanyId

	log.Printf("fetching song by name: %v and companyId: %v", songName, companyId)
	inventoryByName := songsDA.GetSongByName(songName, companyId)
	if inventoryByName != nil && inventoryByName.Id > 0 {
		return 0, errors.New("Song already exists!!")
	}

	// NOTE: ideally insertion into these two tables should be done atomically.

	inventory := getInventory(request)

	err := songsDA.AddSongToInventory(inventory)
	if err != nil {
		return 0, err
	}
	inventoryByNameModel := getInventoryByName(inventory)

	err = songsDA.Create(inventoryByNameModel)
	return inventory.Id, err
}

func getInventoryByName(inventory *songModels.Inventory) *songModels.InventorySortedByName {
	return &songModels.InventorySortedByName{
		Id:        inventory.Id,
		CompanyId: inventory.CompanyId,
		Name:      inventory.Name,
	}
}

func getInventory(request *songs.AddSongToInventory) *songModels.Inventory {
	return &songModels.Inventory{
		CompanyId:  request.CompanyId,
		Id:         idGen.GetIdGenerator().New(),
		Name:       request.Name,
		CreateOn:   time.Now(),
		ModifiedOn: time.Now(),
	}

}
