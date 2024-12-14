package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=username dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
