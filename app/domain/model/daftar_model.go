package model

type DaftarModel struct {
	Name        string `json:"nama"`
	DateOfBirth string `json:"tgl_lahir"`
	Gender      string `json:"gender"`
	NoHp        string `json:"no_hp"`
	NoId        string `json:"no_id"`
	Email       string `json:"email"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}
