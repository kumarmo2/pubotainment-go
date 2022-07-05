package authentication

type SignInRequest struct {
	CompanyId int64  `json:"companyId"`
	Password  string `json:"password"`
}
