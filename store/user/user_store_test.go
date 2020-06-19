package user

import (
	"testing"
)

// TestUsers
func TestUsers(t *testing.T) {
	users, err := Users()
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(users[0].Id)
	for _, user := range users {
		t.Log(user)
	}
}
