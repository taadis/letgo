package store

import (
	"database/sql"
	"log/slog"
)

type PgConfig struct {
	DSN string
	//user=postgres password=postgres dbname=postgres sslmode=disable
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Sslmode  string `json:"sslmode"`
	// todo: 添加其他配置
}

func NewDB(cfg *PgConfig) *sql.DB {
	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		slog.Error(err.Error())
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	//db.SetConnMaxLifetime(time.Hour)
	return db
}
