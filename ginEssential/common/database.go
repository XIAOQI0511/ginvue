package common

import (
	"github.com/jinzhu/gorm"
	"qi.xiao/ginessential/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "ginessential.db")
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	// defer db.Close()
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
