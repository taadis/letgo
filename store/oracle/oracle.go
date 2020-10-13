package oracle

import (
	"database/sql"
	_ "github.com/godror/godror"
)

var (
	OraDb *sql.DB
)

func init() {
	driverName := "godror"
	// username/password@ip:port/sid
	dataSourceName := "neands5/abc123@10.0.17.70:1521/orcl"
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	OraDb = db
}
