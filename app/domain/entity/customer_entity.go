package entity

type CustomerData struct {
	BaseModel
	Name        string `gorm:"column:nama_customer;size:100;"`
	DateOfBirth string `gorm:"column:date_of_birth;type:date;"`
	Gender      string `gorm:"column:gender;size:2;"`
	NoHp        string `gorm:"column:no_handphone;size:13;"`
	NoId        string `gorm:"column:no_id;size:25;index:idx_unique_data,unique"`
	Email       string `gorm:"column:email;size:25;index:idx_unique_data,unique"`
	Status      int    `gorm:"column:status"`
}
