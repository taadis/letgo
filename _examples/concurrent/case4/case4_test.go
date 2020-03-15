package case4

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

var user *User
var once sync.Once

type User struct{

}

func SingleInstance() *User {
	once.Do(func() {
		fmt.Print("once")
		user = new(User)
	})
	return user
}

func TestSingleInstance(t *testing.T)  {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++{
		wg.Add(1)
		go func() {
			user := SingleInstance()
			t.Log(unsafe.Pointer(user))
			wg.Done()
		}()
	}
	wg.Wait()
}
