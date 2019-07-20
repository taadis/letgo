package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dbContext *sql.DB

//
func init() {
	dbContext, err := sql.Open("sqlite3", "./my.db")
	if err != nil {
		log.Panicln(err)
	}
}\

// 报错如下:
//  exec: "gcc": executable file not found in %PATH%
// 原因是sqlitle3是个cgo库，需要gcd编译c代码
// 然后下载安装tdm-gcc即可（windosw版本）下载地址：http://tdm-gcc.tdragon.net/download

func main() {
	log.Println("test")
}
