package entity

import "time"

type DataOrder struct {
	BaseModel
	IdProduct    string       `gorm:"column:id_produck;size:100"`
	ProductData  ProductData  `gorm:"foreignKey:IdProduct"`
	IdCustomer   string       `gorm:"column:id_customer;size:100"`
	CustomerData CustomerData `gorm:"foreignKey:IdCustomer"`
	OrderDate    time.Time    `gorm:"column:order_date"`
	Qty          int
}
