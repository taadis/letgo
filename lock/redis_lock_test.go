package lock

import (
	"context"
	"testing"
	"time"
)

func TestRedisLock_Lock(t *testing.T) {
	ctx := context.Background()
	l, err := NewRedisLock(ctx, "test20220124", time.Second*3)
	if err != nil {
		t.Fatal(err)
	}

	result, err := l.Lock(ctx)
	if err != nil {
		t.Fatalf("lock error:%+v", err)
	}
	t.Logf("lock result:%v", result)
}

func TestRedisLock_Unlock(t *testing.T) {
	ctx := context.Background()
	l, err := NewRedisLock(ctx, "test:redis:lock", time.Second*3)
	if err != nil {
		t.Fatalf("new redis lock error:%+v", err)
	}

	_, err = l.Lock(ctx)
	if err != nil {
		t.Fatalf("redis lock error:%+v", err)
	}

	_, err = l.Unlock(ctx)
	if err != nil {
		t.Fatalf("redis unlock error:%+v", err)
	}
}
