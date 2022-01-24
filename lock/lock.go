package lock

type Locker interface {
	Lock() (bool, error)   // 上锁
	Unlock() (bool, error) // 解锁,可以配置defer使用
	WaitLock() error       // 等待上锁,抢占式上锁,自旋锁
}
