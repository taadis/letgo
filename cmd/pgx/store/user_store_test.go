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

	connPool := NewDBPool(cfg)
	defer Close()

	userStore := NewUserStore(connPool)

	t.Run("TestQueryUser", func(t *testing.T) {
		rets, err := userStore.QueryData(ctx)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("rets: %+v", rets)
	})

	t.Run("TestDeleteUser", func(t *testing.T) {
		err := userStore.DeleteData(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("delete success")
	})

	// todo:more test...
}
