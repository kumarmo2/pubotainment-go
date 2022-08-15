package songs

import (
	// - "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	dataaccess "pubwebservice/dataAccess"
	"pubwebservice/models/songs"
)

func GetSongByName(name string, companyId int64) *songs.InventorySortedByName {
	conn, err := dataaccess.GetCassandraConnectionFactory().GetConnection()

	if err != nil {
		panic(err.Error())
	}
	var cid int64
	var n string
	var id int64

	err = conn.Query("select companyid, name, id  from songs.inventory_sorted_by_name where companyid = ? and name = ?",
		companyId, name).Scan(&cid, &n, &id)

	if err != nil {
		log.Printf("err: %v", err.Error())
	}
	return &songs.InventorySortedByName{
		CompanyId: cid,
		Name:      n,
		Id:        id,
	}
}

func Create(inventory *songs.InventorySortedByName) error {
	conn, err := dataaccess.GetCassandraConnectionFactory().GetConnection()

	if err != nil {
		panic(err.Error())
	}

	err = conn.Query(`insert into songs.inventory_sorted_by_name(companyid , id , name)
        values( ?, ?, ? )`, inventory.CompanyId, inventory.Id, inventory.Name).Exec()

	return err
}
