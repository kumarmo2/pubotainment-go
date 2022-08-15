package songs

type InventorySortedByName struct {
	CompanyId int64  `cql:"companyid"`
	Name      string `cql:"name"`
	Id        int64  `cql:"id"`
}
