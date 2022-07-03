package authentication

import (
	"fmt"
	authModels "pubwebservice/models/authentication"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetCompanyByName(name string) *authModels.Company {
	db, err := sqlx.Connect("postgres", "user=postgres password=admin dbname=pubotainment sslmode=disable")

	if err != nil {
		panic("could not connect to db")
	}

	var company authModels.Company
	fmt.Printf("name: %v", name)
	err = db.Get(&company, "select * from authentication.companies where name=$1", name)
	if err != nil {
		panic(fmt.Sprintf("error while fetching company by name, err: %v", err))
	}

	return &company
}

func UpdateAdminPass(companyId int64, hashedPass string) error {

	db, err := sqlx.Connect("postgres", "user=postgres password=admin dbname=pubotainment sslmode=disable")
	if err != nil {
		panic("could not connect to db")
	}

	fmt.Printf("\n\nid: %v,\n\nhashedPass: %v\n\n", companyId, hashedPass)

	result, err := db.NamedExec(`update authentication.companies set adminhashedpass='hashedpass' where id = 1`,
		map[string]interface{}{
			"hashedpass": "sdfsdfk",
			"id":         companyId,
		})
	fmt.Printf("sdfsdlkf lsdfkls slflksdf \n\n")
	if err != nil {
		panic(fmt.Sprintf("err: %v", err.Error()))
	}
	rows, err := result.RowsAffected()

	if err != nil {
		panic(fmt.Sprintf("err: %v", err.Error()))
	}
	// rowId, _ := result.LastInsertId()
	// fmt.Printf("rowId: %v\n rows: %v\n\n", rowId, rows)
	fmt.Printf("rows: %v\n\n", rows)
	// if err != nil {
	// fmt.Printf("\nerr while updating the db, error: %v\n", err)
	// }
	return err

}
