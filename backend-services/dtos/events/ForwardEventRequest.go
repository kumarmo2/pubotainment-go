package events

type ForwardEventRequest struct {
	CompanyId int64       `json:"companyId"`
	Payload   interface{} `json:"payload"`
}
