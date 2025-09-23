package models

type Users struct {
	Id       int    `json:"id" gorm:"id"`
	UserName string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	Email    string `json:"email" gorm:"email"`
}

func (Users) TableName() string {
	return "users"
}
