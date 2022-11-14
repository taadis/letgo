package cache

import (
	"context"
	"errors"
	"time"
)

var (
	// DefaultCache is the default cache.
	DefaultCache Cache = NewCache()

	// DefaultExpiration is the default duration for items stored in the cache to expire.
	DefaultExpiration time.Duration = 0

	ErrItemExpired = errors.New("item has expired")

	ErrKeyNotFound = errors.New("key not found in cache")
)

// Cache is the interface that wraps the cache.
type Cache interface {
	Get(ctx context.Context, key string) (interface{}, time.Time, error)
	Set(ctx context.Context, key string, value interface{}, d time.Duration) error
	Del(ctx context.Context, key string) error
}

// NewCache returns a new cache.
func NewCache(opts ...Option) Cache {
	options := NewOptions(opts...)
	items := make(map[string]Item)

	if len(options.Items) > 0 {
		items = options.Items
	}

	return &memoryCache{
		opts:  options,
		items: items,
	}
}
