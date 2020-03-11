package concurrent

import (
	"testing"
)

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

// 测试只读 channel
func TestReadonlyChannel(t *testing.T) {
	var readonlyChannel <-chan int
	t.Logf("readonlyChannel type is %T", readonlyChannel)
	t.Log(readonlyChannel) // nil
	//t.Log(<-readonlyChannel) // fatal error: all goroutines are asleep - deadlock!
	//readonlyChannel <- 1 // invalid operation: readonlyChannel <- 1 (send to receive-only type <-chan int)
}

// 测试只写 channel
func TestWirteonlyChannel(t *testing.T) {
	var writeonlyChannel chan<- int
	t.Logf("writeonlyChannel type is %T", writeonlyChannel)
	t.Log(writeonlyChannel)
	//<-writeonlyChannel // invalid operation: <-writeonlyChannel (receive from send-only type chan<- int)
	//writeonlyChannel <- 1 // fatal error: all goroutines are asleep - deadlock!
}

// test ReadWriteChannel
func TestReadWriteChannel(t *testing.T) {
	var readWriteChannel chan int
	t.Logf("readWriteChannel type is %T", readWriteChannel)
	t.Log(readWriteChannel)
	ch := make(chan int)
	t.Log(ch)
}

// 使用 make 声明管道
func TestMakeChannel(t *testing.T) {
	// 使用 make 分配内存空间, 不然使用时直接死锁 deadlock
	ch2 := make(chan int)
	t.Log(ch2)
}

// 测试从关闭的管道中读取数据
// 水龙头关了, 能接水(能读取), 但是只能接到空气(只能读取到数据类型的零值).
func TestReadForClosedChannel(t *testing.T) {
	// chan int
	ch_int := make(chan int)
	close(ch_int)
	t.Log(<-ch_int)

	// chan string
	ch_string := make(chan string)
	close(ch_string)
	t.Log(<-ch_string)

	// chan struct{}
	ch_struct := make(chan struct{})
	close(ch_struct)
	t.Log(<-ch_struct)

	// chan interface{}
	ch_interface := make(chan interface{})
	close(ch_interface)
	t.Log(<-ch_interface)

	// chan func{}
	// 这个怎么写?
}
