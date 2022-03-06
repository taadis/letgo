# Go并发实践之生成器

在实际的编程中，往往需要一个统一的编号生成器，比如唯一序列号、订单号、快递单号、随机数等等。

这种生成器需要满足高性能、高并发等要求。用Go来实现此类需求非常合适。

下面来实现一个随机数生成器来说明。

## 初级版本的生成器（带缓冲）

代码实现如下：

```go
package main

import (
	"fmt"
	"math/rand"
)

func GenerateRandomA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func main() {
	ch := GenerateRandomA()
	fmt.Printf("%d\n", <-ch)
	fmt.Printf("%d\n", <-ch)
}

```

在当前目录执行以下命令来运行程序。

```
go run main.go
```

查看输出结果，输出了2行，每行有一个随机数。

```
5577006791947779410
8674665223082153551
```

## 多随机源的生成器

```
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			ch <- rand.Int()
		}
	}()
	return ch
}

// GenerateRandomB 生成1000以内的随机数，以做区分
func GenerateRandomB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			ch <- rand.Intn(1000)
		}
	}()
	return ch
}

func GenerateRandom() chan int {
	ch := make(chan int, 20)
	go func() {
		for {
			// 使用select的fan-in来增加生成器的随机源
			select {
			case ch <- <-GenerateRandomA():
			case ch <- <-GenerateRandomB():
			}
		}
	}()
	return ch
}

func main() {
	ch := GenerateRandom()
	for random := range ch {
		fmt.Printf("%d\n", random)
	}
}

```

在当前目录执行以下命令来运行程序。

```
go run main.go
```

查看输出结果，会发现交错输出GenerateRandomA()生成的大随机数和GenerateRandom生成的小随机数。

```
3583801852838737519
735
4374228143828497897
2513535816493232224
517
197
124672858875651291
583292530287721155
8895284326904277367
122
```

同时会发现输出停不下来了，可以用快捷键`Ctrl+C`来取消。
