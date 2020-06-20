package store

import (
	"testing"
)

func TestUsers(t *testing.T) {
	var user SystemUser
	Db.First(&user)
	t.Log(user)
}
