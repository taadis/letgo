package user

import (
	"testing"
)

// TestFind
func TestFind(t *testing.T) {
	id := "002f2638-b3f0-4154-9aaa-252b405adc87"
	user, err := Find(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

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
