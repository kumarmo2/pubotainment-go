package servicediscovery

import (
	"pubwebservice/dataAccess"
	"pubwebservice/models/serviceDiscovery"
)

func InsertConnectionMap(connectionMap *servicediscovery.ConnectionServerMap) error {
	conn, err := dataaccess.GetCassandraConnectionFactory().GetConnection()
	if err != nil {
		panic(err.Error())
	}
	err = conn.Query(`insert into servicediscovery.connectionservermap(connectionid, serverid, lastpinged, companyid ) 
                values (?, ?, ?, ?)`, connectionMap.ConnectionId, connectionMap.ServerId,
		connectionMap.LastPinged, connectionMap.CompanyId).Exec()

	if err != nil {
		panic(err.Error())
	}
	return nil
}
