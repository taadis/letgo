package cache

import (
	"context"
	"testing"
)

var (
	ctx               = context.Background()
	key               = "key:test"
	value interface{} = "value:test"
)

func TestMemoryCache(t *testing.T) {
	t.Run("CacheGetMiss", func(t *testing.T) {
		_, _, err := NewCache().Get(ctx, key)
		if err == nil {
			t.Error("")
		}
	})

	t.Run("CacheGetHit", func(t *testing.T) {
		c := NewCache()

		err := c.Set(ctx, key, value, 0)
		if err != nil {
			t.Error(err)
		}

		i, _, err := c.Get(ctx, key)
		if err != nil {
			t.Errorf("expected get value, got error:%+v", err)
		}
		if i != value {
			t.Errorf("expected %s, got %v", value, i)
		}
	})
}
