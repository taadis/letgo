package store

import (
	"testing"
)

// TestPing
func TestPing(t *testing.T) {
	err := Db.Ping()
	if err != nil {
		t.Fatal(err)
	}
	defer Db.Close()
}
