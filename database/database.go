package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"tiktok_info/entity"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	const MYSQL = "root:@tcp(127.0.0.1:3306)/tiktok_user_info?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Database connected")
}

func RunMigration() {
	err := DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database migrated")
}
