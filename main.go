package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type News struct {
	Id uint `gorm:"primaryKey" json:"id" form:"id"`
	Title string `gorm:"not null" json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main(){
	InitDatabase()
	e := echo.New()
	e.POST("/news", CreateNewsController)
	e.GET("/news", GetNewsController)
	e.Start(":8000")
}

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja12?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
        panic("failed to connect database")
    }
	InitMigration()
}

func InitMigration(){
	DB.AutoMigrate(&News{})
}

func CreateNewsController(c echo.Context) error {
	var newsRequest News
	c.Bind(&newsRequest)

	result := DB.Create(&newsRequest)
	if result.Error != nil {
        return c.JSON(500, result.Error.Error())
    }
	return c.JSON(200, newsRequest)
} 

func GetNewsController(c echo.Context) error {
	var news []News
	result := DB.Find(&news)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	return c.JSON(http.StatusOK, news)
}



