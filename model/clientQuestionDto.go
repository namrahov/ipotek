package model

type ClientQuestionRequest struct {
	tableName struct{} `sql:"client_question_request" pg:",discard_unknown_columns"`

	Id        int64  `sql:"id"  json:"id"`
	ProductId int64  `sql:"product_id" json:"productId"`
	Name      string `sql:"name" json:"name"`
	Surname   string `sql:"surname" json:"surname"`
	Phone     string `sql:"phone" json:"phone"`
	CreatedAt string `sql:"created_at" json:"createdAt"`
	UpdatedAt string `sql:"updated_at" json:"updatedAt"`
}

type ClientRequest struct {
	tableName struct{} `sql:"client_request" pg:",discard_unknown_columns"`

	Id          int64  `sql:"id"  json:"id"`
	ProductId   int64  `sql:"product_id" json:"productId"`
	Name        string `sql:"name" json:"name"`
	Surname     string `sql:"surname" json:"surname"`
	PinCode     string `sql:"pin_code" json:"pinCode"`
	Phone       string `sql:"phone" json:"phone"`
	DateOfBirth string `sql:"date_of_birth" json:"dateOfBirth"`
	CreatedAt   string `sql:"created_at" json:"createdAt"`
	UpdatedAt   string `sql:"updated_at" json:"updatedAt"`
}

type ClientAgreement struct {
	tableName struct{} `sql:"client_agreement" pg:",discard_unknown_columns"`

	Id            int64  `sql:"id"  json:"id"`
	ClientId      int64  `sql:"client_id" json:"clientId"`
	AgreementText string `sql:"agreement_text" json:"agreementText"`
	IsSigned      bool   `sql:"is_signed" json:"isSigned"`
	SignDate      string `sql:"sign_date" json:"signDate"`
}

type ClientQuestionDto struct {
	Name        string      `json:"name"`
	ProductId   int64       `json:"productId"`
	Surname     string      `json:"surname"`
	Phone       string      `json:"phone"`
	ProductName ProductName `json:"productName"`
}

type ClientDto struct {
	Name               *string             `json:"name"`
	ProductId          int64               `json:"productId"`
	Surname            *string             `json:"surname"`
	Phone              string              `json:"phone"`
	PinCode            string              `json:"pinCode"`
	DateOfBirth        *string             `json:"dateOfBirth"`
	ClientAgreementDto *ClientAgreementDto `json:"clientAgreementDto"`
	ProductName        ProductName         `json:"productName"`
}

type ClientAgreementDto struct {
	AgreementText string `json:"agreementText"`
	IsSigned      bool   `json:"isSigned"`
	SignDate      string `json:"signDate"`
}

type ProductName string

const (
	ExtractiveDomesticMortgage ProductName = "ExtractiveDomesticMortgage"
	UnsecuredDomesticMortgage  ProductName = "UnsecuredDomesticMortgage"
	MidaMortgage               ProductName = "MidaMortgage"
	StateMortgage              ProductName = "StateMortgage"
	PreferentialStateMortgage  ProductName = "PreferentialStateMortgage"
	PSQ                        ProductName = "PSQ"
)
