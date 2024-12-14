package store

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
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

var connPool *pgxpool.Pool

// 引入 sync.Once 来确保 NewDB 函数中的连接池代码只执行一次
// 使用 once.Do 方法包裹数据库连接池的初始化逻辑，这样即使 NewDB 被多协程或并发多次调用，也只会在第一次调用时初始化连接池。
var once sync.Once

func NewDBPool(cfg *PgConfig) *pgxpool.Pool {
	once.Do(func() {
		var err error
		connPool, err = pgxpool.New(context.Background(), cfg.DSN)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
			os.Exit(1)
		}

		// 测试连接是否成功
		ctx := context.Background()
		err = connPool.Ping(ctx)
		if err != nil {
			slog.Error("ping db failed", "error", err)
			os.Exit(1)
		}
		slog.InfoContext(ctx, "ping db success")
	})

	return connPool
}

func Close() {
	// 引入 sync.Once 来确保 Close 函数中的连接池关闭代码只执行一次
	slog.Info("close db pool...")
	once.Do(func() {
		if connPool == nil {
			slog.Info("db pool is nil not need to close")
			return
		}

		connPool.Close()
		slog.Info("closed db pool")
	})
}
