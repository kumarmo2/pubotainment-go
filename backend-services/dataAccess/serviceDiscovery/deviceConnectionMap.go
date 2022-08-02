package servicediscovery

import (
	"pubwebservice/dataAccess"
	"pubwebservice/models/serviceDiscovery"
)

func InsertConnectionMap(connectionMap *servicediscovery.DeviceConnectionMap) error {
	conn, err := dataaccess.GetCassandraConnectionFactory().GetConnection()
	if err != nil {
		panic(err.Error())
	}
	err = conn.Query(`insert into servicediscovery.deviceconnectionmap(deviceid, serverid, lastpinged, companyid ) 
                values (?, ?, ?, ?)`, connectionMap.DeviceId, connectionMap.ServerId,
		connectionMap.LastPinged, connectionMap.CompanyId).Exec()

	if err != nil {
		panic(err.Error())
	}
	return nil
}
