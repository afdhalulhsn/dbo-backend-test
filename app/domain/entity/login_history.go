package entity

type LoginHistory struct {
	BaseModel
	IdUser      string   `gorm:"column:id_user;size:100"`
	UserData    UserData `gorm:"foreignKey:IdUser"`
	LoginTime   string   `gorm:"column:login_time;type:timestamp;"`
	ExpiredTime string   `gorm:"column:expired_time;type:timestamp;"`
	Status      string   `gorm:"column:status;size:2"`
}
