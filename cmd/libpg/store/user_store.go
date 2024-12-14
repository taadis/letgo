package store

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserStore interface {
	InsertData(ctx context.Context) error
	QueryData(ctx context.Context) ([]User, error)
	UpdateData(ctx context.Context) error
	DeleteData(ctx context.Context) error
}

// 导入配置对象，进行db初始化
func NewUserStore(cfg *PgConfig) UserStore {
	s := &UserStoreImpl{}
	s.db = NewDB(cfg)
	return s
}

type UserStoreImpl struct {
	db *sql.DB
}

func (s *UserStoreImpl) QueryData(ctx context.Context) ([]User, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users, nil
}
func (s *UserStoreImpl) InsertData(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO users (name, age) VALUES ($1, $2)", "Alice", 30)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *UserStoreImpl) UpdateData(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, "UPDATE users SET age = $1 WHERE name = $2", 31, "Alice")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *UserStoreImpl) DeleteData(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM users WHERE name = $1", "Alice")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
