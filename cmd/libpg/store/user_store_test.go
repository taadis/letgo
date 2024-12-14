package store

import (
	"context"
	"testing"
)

func TestUserStore(t *testing.T) {
	ctx := context.Background()
	cfg := &PgConfig{}
	cfg.DSN = "service=dev"
	//cfg.DSN = "user=postgres password=postgres dbname=postgres sslmode=disable",

	// db, err := sql.Open("postgres", cfg.DSN)
	// if err != nil {
	// 	t.Fatalf("Failed to connect to database: %v", err)
	// }
	// defer db.Close()

	userStore := NewUserStore(cfg)

	t.Run("TestQueryUser", func(t *testing.T) {
		rets, err := userStore.QueryData(ctx)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("rets: %+v", rets)
	})

	//userStore.InsertData(db)
	//userStore.QueryData(db)
	//userStore.UpdateData(db)
	//userStore.DeleteData(db)
}
