package store

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStore interface {
	InsertData(ctx context.Context) error
	QueryData(ctx context.Context) ([]*User, error)
	UpdateData(ctx context.Context) error
	DeleteData(ctx context.Context, id int64) error
}

// 导入配置对象，进行db初始化
func NewUserStore(connPool *pgxpool.Pool) UserStore {
	s := &UserStoreImpl{}
	s.db = connPool
	return s
}

type UserStoreImpl struct {
	db *pgxpool.Pool
}

func (s *UserStoreImpl) QueryData(ctx context.Context) ([]*User, error) {
	rows, err := s.db.Query(ctx, "SELECT id, name, age FROM users")
	if err != nil {
		slog.ErrorContext(ctx, "db query", "error", err)
		return nil, err

	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			slog.ErrorContext(ctx, "db scan", "error", err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
func (s *UserStoreImpl) InsertData(ctx context.Context) error {
	_, err := s.db.Exec(ctx, "INSERT INTO users (name, age) VALUES ($1, $2)", "Alice", 30)
	if err != nil {
		slog.ErrorContext(ctx, "db exec", "error", err)
		return err
	}

	return nil
}

func (s *UserStoreImpl) UpdateData(ctx context.Context) error {
	_, err := s.db.Exec(ctx, "UPDATE users SET age = $1 WHERE name = $2", 31, "Alice")
	if err != nil {
		slog.ErrorContext(ctx, "db exec", "error", err)
		return err
	}

	return nil
}

func (s *UserStoreImpl) DeleteData(ctx context.Context, id int64) error {
	_, err := s.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		slog.ErrorContext(ctx, "db exec", "error", err)
		return err
	}

	return nil
}
