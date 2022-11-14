package cache

import (
	"context"
	"sync"
	"time"
)

type memoryCache struct {
	opts Options
	sync.RWMutex

	items map[string]Item
}

func (c *memoryCache) Get(ctx context.Context, key string) (interface{}, time.Time, error) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, time.Time{}, ErrKeyNotFound
	}

	if item.Expired() {
		return nil, time.Time{}, ErrItemExpired
	}

	return item.Value, time.Unix(0, item.Expiration), nil
}

func (c *memoryCache) Set(ctx context.Context, key string, value interface{}, d time.Duration) error {
	var expiration int64
	if d == DefaultExpiration {
		d = c.opts.Expiration
	}
	if d > 0 {
		expiration = time.Now().Add(d).UnixNano()
	}

	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
	}

	return nil
}

func (c *memoryCache) Del(ctx context.Context, key string) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	_, found := c.items[key]
	if !found {
		return ErrKeyNotFound
	}

	delete(c.items, key)
	return nil
}
