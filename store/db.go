package store

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/taadis/letgo/conf"
	// "github.com/spf13/viper"
)

var Db *gorm.DB

// init
func init() {
	database := conf.Database()
	driverName := database.DriverName
	dataSourceName := database.DataSourceName
	logModel := database.LogMode
	db, err := gorm.Open(driverName, dataSourceName)
	db.LogMode(logModel)
	if err != nil {
		log.Fatalln("gorm.Open error", err.Error())
	}
	Db = db
}
