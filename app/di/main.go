package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/taadis/letgo/app/di/service"
	"github.com/taadis/letgo/app/di/store"
)

func main() {
	// 这里使用mysql的驱动实现,可以自行替换成其他pg等
	dsn := "user:password@tcp(127.0.0.1:3306)/database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open db error:%+v", err)
	}

	// 创建store实例,内部依赖项为db
	dbStore := store.NewStore(db)
	// 通过注入store作为依赖项来创建service的实例srv
	srv := &service.Service{Store: dbStore}

	// 下面的代码实现了一个简单的命令行应用,以读取id作为输入,从数据库存储查询相关id的结果
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Print(err)
			continue
		}
		result, err := srv.GetSome(id)
		if err != nil {
			fmt.Printf("get some error:%+v", err)
			continue
		}
		fmt.Printf("get some result:%d", result)
	}
}
