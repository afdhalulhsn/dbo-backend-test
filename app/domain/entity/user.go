package entity

type UserData struct {
	BaseModel
	UserName     string       `gorm:"column:username;size:25;unique"`
	Password     string       `gorm:"column:password;size:25;,"`
	Role         string       `gorm:"column:role;size:25;"`
	IdCustomer   string       `gorm:"column:id_customer;size:100"`
	CustomerData CustomerData `gorm:"foreignKey:IdCustomer"`
	Token        string       `gorm:"-"`
	Status       int          `gorm:"column:status"`
}
