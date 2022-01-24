package lock

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisLock struct {
	key        string
	expiration time.Duration
	client     *redis.Client
}

func NewRedisLock(ctx context.Context, key string, expiration time.Duration) (*RedisLock, error) {
	opts := &redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(opts)
	client.WithContext(ctx)
	return &RedisLock{
		key:        key,
		expiration: expiration,
		client:     client,
	}, nil
}

func (l *RedisLock) Lock(ctx context.Context) (bool, error) {
	const value byte = 1
	return l.client.SetNX(ctx, l.key, value, l.expiration).Result()
}

func (l *RedisLock) Unlock(ctx context.Context) (bool, error) {
	result, err := l.client.Del(ctx, l.key).Result()
	if err != nil {
		return false, err
	}

	if result == 0 {
		return false, err
	}

	return true, nil
}

func (l *RedisLock) WaitLock(ctx context.Context) error {
	for {
		result, err := l.client.Exists(ctx, l.key).Result()
		if err != nil {
			return err
		}

		if result == 0 {
			locked, err := l.Lock(ctx)
			if err != nil {
				return err
			}
			if locked {
				return nil
			}
		}
		time.Sleep(time.Millisecond * 10)
	}
}
