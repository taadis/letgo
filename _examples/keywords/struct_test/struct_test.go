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

// 声明一个结构体的行为, 也就是其下的函数
func (user *User) Foo() {

}

// 声明一个结构体的行为,
// 不使用指针时, 会复制一份实例对象及其成员, 会造成额外的性能开销,
// 所以通常生命结构体的行为, 也就是函数时,  更多的是使用结构体的指针
func (user User) Foo1() {

}
