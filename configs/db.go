package configs

import (
	"prakerja12/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


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
	DB.AutoMigrate(&models.News{}, &models.User{})
}