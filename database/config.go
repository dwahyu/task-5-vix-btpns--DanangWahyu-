package database

import (
	"fmt"

	"task-5-vix-btpns/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	db, err := gorm.Open(mysql.Open("dan:kosong@tcp(127.0.0.1:3308)/task_btpns_api"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	db.AutoMigrate(&models.User{})

	DB = db
}