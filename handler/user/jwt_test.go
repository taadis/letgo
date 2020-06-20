package user

import (
	"testing"
)

// TestGenerateToken
func TestGenerateToken(t *testing.T) {
	token, err := GenereateToken("xxx")
	if err != nil {
		t.Fatal("generate token error", err.Error())
	}
	t.Log("generated token", token)
}
