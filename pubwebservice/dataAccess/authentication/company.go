package authentication

import (
	"fmt"
	authModels "pubwebservice/models/authentication"

	// "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	da "pubwebservice/dataAccess"
)

func GetCompanyByName(name string) *authModels.Company {
	db, err := da.GetConnectionFactory().GetConnection()

	if err != nil {
		panic("could not connect to db")
	}

	var company authModels.Company

	err = db.Get(&company, "select * from authentication.companies where name=$1", name)
	if err != nil {
		panic(fmt.Sprintf("error while fetching company by name, err: %v", err))
	}

	return &company
}

func UpdateUserPass(companyId int64, hashedPass string) error {
	db, err := da.GetConnectionFactory().GetConnection()
	if err != nil {
		panic("could not connect to db")
	}

	_, err = db.NamedExec(`update authentication.companies set userhashedpass=:hashedpass where id =:id`,
		map[string]interface{}{
			"hashedpass": hashedPass,
			"id":         companyId,
		})

	if err != nil {
		panic(fmt.Sprintf("err: %v", err.Error()))
	}
	return nil
}

func UpdateAdminPass(companyId int64, hashedPass string) error {

	db, err := da.GetConnectionFactory().GetConnection()
	if err != nil {
		panic("could not connect to db")
	}
	_, err = db.NamedExec(`update authentication.companies set adminhashedpass=:hashedpass where id =:id`,
		map[string]interface{}{
			"hashedpass": hashedPass,
			"id":         companyId,
		})

	if err != nil {
		panic(fmt.Sprintf("err: %v", err.Error()))
	}
	return nil
}
