package authentication

import (
	"errors"
	"fmt"
	"log"
	authDA "pubwebservice/dataAccess/authentication"
	authDtos "pubwebservice/dtos/authentication"

	"golang.org/x/crypto/bcrypt"
)

func SignInAdmin(request *authDtos.SignInRequest) bool {
	if request == nil {
		panic("signIn request cannot be null.")
	}
	company := authDA.GetCompanyById(request.CompanyId)

	if company == nil || company.Id < 1 {
		log.Println("company not found. companyId:", request.CompanyId)
		return false
	}
	companyAdminHash := company.AdminHashedPass
	pass := request.Password

	if err := bcrypt.CompareHashAndPassword([]byte(*companyAdminHash), []byte(pass)); err != nil {
		log.Println("error while comparing admin's pass's hash", err.Error())
		return false
	}
	return true
}

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

// func IsValidAdminSignInRequest(request *authDtos.SignInRequest) (bool, error) {
// if request = nil {
// return false, errors.New("Invalid signIn Request")
// }

// }
