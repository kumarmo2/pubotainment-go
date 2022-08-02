package servicediscovery

import "time"

type DeviceConnectionMap struct {
	ServerId   string    `cql:"serverid"`
	DeviceId   string    `cql:"deviceid"`
	LastPinged time.Time `cql:"lastpinged"`
	CompanyId  int64     `cql:"companyid"`
}
