package store

import (
	"database/sql"
	"log"

	"gitee.com/taadis/letgo/conf"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

// init
func init() {
	database := conf.Database()
	driverName := database.DriverName
	dataSourceName := database.DataSourceName
	Db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalln("sql open error", err.Error())
	}
}
