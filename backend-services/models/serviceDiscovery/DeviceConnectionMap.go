package servicediscovery

import "time"

type ConnectionServerMap struct {
	ServerId     string    `cql:"serverid"`
	ConnectionId string    `cql:"connectionid"`
	LastPinged   time.Time `cql:"lastpinged"`
	CompanyId    int64     `cql:"companyid"`
}
