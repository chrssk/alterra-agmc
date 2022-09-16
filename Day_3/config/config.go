package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"Day2Assignment/models"
)

var (
	DB *gorm.DB
)

func InitDB(){

	connectionString := "root:@tcp(localhost:3306)/AdvancedGolang?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil{
		panic(err)
	}
}

func InitialMigration(){
	DB.AutoMigrate(&models.User{})
}