package store

import (
	"log"

	"gitee.com/taadis/letgo/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// "github.com/spf13/viper"
)

var Db *gorm.DB

// init
func init() {
	database := conf.Database()
	driverName := database.DriverName
	dataSourceName := database.DataSourceName
	db, err := gorm.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalln("gorm.Open error", err.Error())
	}
	Db = db
}
