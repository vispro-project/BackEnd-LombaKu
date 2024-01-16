package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db_lombaku"))
	if err != nil {
		panic(err)
	}
	// AutoMigrate untuk tabel Lomba
	err = db.AutoMigrate(&Lomba{})
	if err != nil {
		panic(err)
	}

	// AutoMigrate untuk tabel User
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	DB = db

}
