package songs

import "time"

type Inventory struct {
	CompanyId  int64     `cql:"companyid"`
	Id         int64     `cql:"id"`
	Name       string    `cql:"name"`
	CreateOn   time.Time `cql:"createon"`
	ModifiedOn time.Time `cql:"modifiedon"`
}
