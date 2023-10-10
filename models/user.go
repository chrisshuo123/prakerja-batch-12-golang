package models

type User struct {
	Id uint `gorm:"primaryKey" json:"id" form:"id"`
	Name string `gorm:"not null" json:"name"`
	Email string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type UserResponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}