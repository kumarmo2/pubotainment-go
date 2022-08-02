package servicediscovery

import (
	"log"
	"pubwebservice/dataAccess"
	"time"
)

func InsertServcerInstance(serverId string, ips []string) {
	conn, err := dataaccess.GetCassandraConnectionFactory().GetConnection()
	if err != nil {
		panic(err.Error())
	}
	log.Println("========= inserting ======")
	err = conn.Query("insert into servicediscovery.serverinstances (id, ips, registeredon) values (?, ?, ?)", serverId, ips, time.Now()).Exec()

	if err != nil {
		panic(err.Error())
	}

}

func GetIps(serverId string) []string {
	conn, err := dataaccess.GetCassandraConnectionFactory().GetConnection()
	if err != nil {
		panic(err.Error())
	}
	var ips []string
	log.Println("making query")

	err = conn.Query("select ips from servicediscovery.serverinstances where id = ?", serverId).Scan(&ips)
	if err != nil {
		log.Println("err:", err.Error())
	}
	return ips
}
