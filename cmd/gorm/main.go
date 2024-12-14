package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	//dsn := "host=my-postgres.orb.local user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	// 使用服务名连接，将从~/.pg_service.conf中读取连接信息
	dsn := "service=dev"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})

	// 创建
	db.Create(&User{Name: "Alice", Email: "alice@example.com"})

	// 读取
	var user User
	db.First(&user, 1)                                // 根据整型主键查找
	db.First(&user, "email = ?", "alice@example.com") // 查找email为alice@example.com的用户

	// 更新 - 更新用户的email
	db.Model(&user).Update("Email", "alice@newdomain.com")

	// 删除 - 删除用户
	db.Delete(&user, 1)
}
