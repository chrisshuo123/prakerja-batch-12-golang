package models

type News struct {
	Id uint `gorm:"primaryKey" json:"id" form:"id"`
	Title string `gorm:"not null" json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
}