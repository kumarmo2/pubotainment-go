package authentication

import (
	"errors"
	"fmt"
	authDA "pubwebservice/dataAccess/authentication"
	authDtos "pubwebservice/dtos/authentication"

	"golang.org/x/crypto/bcrypt"
)

func X() {}

func RegisterUser(companyName string, request authDtos.UserRegistrationRequest) error {
	company := authDA.GetCompanyByName(companyName)

	if company == nil {
		panic(fmt.Sprintf("no company with name: %v found", companyName))
	}
	if company.UserHashedPass != nil {
		return errors.New("User is already registered")
	}
	bytes := []byte(request.Password)

	hashedPassBytes, err := bcrypt.GenerateFromPassword(bytes, 12)

	if err != nil {
		fmt.Printf("error while generating password, err: %v", err.Error())
		panic("Internal Server Error. Please try again")
	}

	hashedPass := string(hashedPassBytes)
	authDA.UpdateUserPass(company.Id, hashedPass)

	return nil
}

func RegisterAdmin(companyName string, request authDtos.AdminRegistrationRequest) error {
	company := authDA.GetCompanyByName(companyName)

	if company == nil {
		panic(fmt.Sprintf("no company with name: %v found", companyName))
	}
	fmt.Printf("company: %+v", company)

	if company.AdminHashedPass != nil {
		return errors.New("Admin is already registered")
	}
	bytes := []byte(request.Password)

	hashedPassBytes, err := bcrypt.GenerateFromPassword(bytes, 12)
	if err != nil {
		fmt.Printf("error while generating password, err: %v", err.Error())
		panic("Internal Server Error. Please try again")
	}

	hashedPass := string(hashedPassBytes)
	authDA.UpdateAdminPass(company.Id, hashedPass)

	return nil
}
