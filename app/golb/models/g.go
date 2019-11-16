package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func ConnectToDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/fenlei?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect database")
	}
	db.SingularTable(true)
	return db
}
