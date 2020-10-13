package oracle

import (
	"testing"
)

func TestOraclePing(t *testing.T) {
	defer OraDb.Close()
	err := OraDb.Ping()
	if err != nil {
		t.Error(err)
	}
}
