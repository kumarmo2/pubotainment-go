package songs

import (
	"pubwebservice/dataAccess"
	// songDtos "pubwebservice/dtos/songs"
	songModels "pubwebservice/models/songs"
)

func AddSongToInventory(inventory *songModels.Inventory) error {
	conn, err := dataaccess.GetCassandraConnectionFactory().GetConnection()

	if err != nil {
		panic(err)
	}

	err = conn.Query(`insert into songs.inventory_main(companyid , id , name , createon , modifiedon)
        values( ?, ?, ?, ?, ?)`, inventory.CompanyId, inventory.Id, inventory.Name, inventory.CreateOn, inventory.ModifiedOn).Exec()

	return err
}
