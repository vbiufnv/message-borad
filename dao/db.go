package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"message-borad/model"
)

var DB *gorm.DB

func ConDB() (err error) {

	dsn := "root:147258aa@tcp(127.0.0.1:3306)/dyt?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&model.Comment{}, &model.User{}, &model.Post{}, &model.Like{})

	return DB.DB().Ping()

}
