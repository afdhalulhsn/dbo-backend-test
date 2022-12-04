package model

type UpdateOrder struct {
	IdOrder string `json:"id_order"`
	Qty     int    `json:"qty"`
}

type InsertOrder struct {
	IdProduct  string `json:"id_product"`
	IdCustomer string `json:"id_customer"`
	Qty        int    `json:"qty"`
}
