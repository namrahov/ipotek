package model

type Product struct {
	tableName struct{} `sql:"product" pg:",discard_unknown_columns"`

	Id          int64   `sql:"id"  json:"id"`
	Name        string  `sql:"name" json:"name"`
	NameAz      string  `sql:"name_az" json:"nameAz"`
	Status      Status  `sql:"status" json:"status"`
	Email       *string `sql:"email" json:"email"`
	Description string  `sql:"description" json:"description"`
	CreatedAt   string  `sql:"created_at" json:"createdAt"`
}

type Status string

const (
	Active   Status = "ACTIVE"
	Deactive Status = "DEACTIVE"
)
