package memory

import (
	"context"
	"sync"
	"time"

	"github.com/taadis/letgo/cache"
)

func NewCache() cache.Cache {
	c := new(memoryCache)
	c.items = make(map[string]*Item)
	return c
}

type Item struct {
	Value      interface{}
	Expiration int64
}

// Expired return true if the item has expired.
func (i *Item) Expired() bool {
	if i.Expiration <= 0 {
		return false
	}

	return time.Now().UnixNano() > i.Expiration
}

type memoryCache struct {
	items map[string]*Item
	sync.RWMutex
}

func (c *memoryCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	c.items[key] = &Item{
		Value:      value,
		Expiration: time.Now().Add(expiration).UnixNano(),
	}

	return nil
}

func (c *memoryCache) Get(ctx context.Context, key string) (interface{}, error) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, cache.ErrKeyNotFound
	}

	if item.Expired() {
		return nil, cache.ErrExpired
	}

	return item.Value, nil
}

func (c *memoryCache) Del(ctx context.Context, key string) error {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()

	_, found := c.items[key]
	if !found {
		return cache.ErrKeyNotFound
	}

	delete(c.items, key)
	return nil
}
