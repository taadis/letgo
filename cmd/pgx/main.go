package main

import (
	"log/slog"
	"net/http"

	"github.com/taadis/letgo/cmd/pgx/handler"
	"github.com/taadis/letgo/cmd/pgx/store"
)

func main() {
	cfg := &store.PgConfig{}
	//cfg.DSN = "user=username dbname=testdb sslmode=disable"
	cfg.DSN = "service=dev"
	dbPool := store.NewDBPool(cfg)
	defer store.Close()

	// use http
	userStore := store.NewUserStore(dbPool)
	userHandler := handler.NewHandler(userStore)
	http.HandleFunc("/query", userHandler.QueryHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("listen and serve failed", "error", err)
	}
}
