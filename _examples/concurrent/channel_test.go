package concurrent

import "testing"

// 测试管道的声明
func TestChannel1(t *testing.T) {
	// 只读 channel
	var readonlyChannel <-chan int
	t.Logf("readonlyChannel type is %T", readonlyChannel)

	// 只写 channel
	var writeonlyChannel chan<- int
	t.Logf("writeonlyChannel type is %T", writeonlyChannel)

	// 可读可写 channel
	var readWriteChannel chan int
	t.Logf("readWriteChannel type is %T", readWriteChannel)
}

// 使用 make 声明管道
func TestMakeChannel(t *testing.T) {
	// 使用 make 分配内存空间, 不然使用时直接死锁 deadlock
	ch2 := make(chan int)
	t.Log(ch2)
}
