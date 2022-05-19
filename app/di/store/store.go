package store

import "database/sql"

type Store interface {
	Get(id int) (int, error)
	Ping() error
}

func NewStore(db *sql.DB) Store {
	return &store{db}
}

// 存储的具体实现可能有多种,本例中为sql.DB的实例,通过db来链接和操作数据库
type store struct {
	db *sql.DB
}

func (s *store) Get(id int) (int, error) {
	// 使用 s.db 进行数据库操作
	return 0, nil
}

func (s *store) Ping() error {
	return s.db.Ping()
}
