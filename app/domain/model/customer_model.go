package model

type CustomerUpdateModel struct {
	IdCustomer  string `json:"id_customer"`
	Name        string `json:"nama"`
	DateOfBirth string `json:"ttl"`
	Gender      string `json:"gender"`
	NoHp        string `json:"no_hp"`
	NoId        string `json:"no_id"`
	Email       string `json:"email"`
}
