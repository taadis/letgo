package metadata

import (
	"context"
	"testing"
)

func TestCopy(t *testing.T) {
	md := Metadata{
		"key1": "value1",
		"key2": "value2",
	}

	cmd := Copy(md)

	for k, v := range md {
		cv := cmd[k]
		if cv != v {
			t.Fatalf("want %s but got %s", v, cv)
		}
	}
}

func TestNewContext(t *testing.T) {
	ctx := context.Background()
	ctx = NewContext(ctx, Metadata{
		"key1": "value1",
		"key2": "value2",
	})
	md, ok := FromContext(ctx)
	if !ok {
		t.Fatalf("want true but got %v", ok)
	}
	if len(md) != 2 {
		t.Fatalf("want 2 but got %d", len(md))
	}
	if md["key1"] != "value1" {
		t.Fatalf("want value1 but got %s", md["key1"])
	}

}
