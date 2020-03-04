// 结构体相关测试
package struct_test

import (
	"testing"
)

// 声明一个结构体, 包含一些属性
type User struct {
	Id   string
	Name string
	Age  int
}

// 测试结构体的实例化/初始化
func TestD(t *testing.T) {
	user := User{}
	user1 := User{Id: "xxx", Name: "zhang san"}
	user2 := new(User) // 指向实例的指针, 等同于 user2 := &User{}
	user3 := &User{}
	t.Log(user)
	t.Logf("user1 type is %T", user1)
	t.Logf("user2 type is %T", user2)
	t.Logf("user3 type is %T", user3)
}
