package concurrent

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go HandleRequest(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("cancel()")
	cancel()
	time.Sleep(5 * time.Second)
}

func HandleRequest(ctx context.Context) {
	// go cache
	// go db
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandleRequest Done")
			return
		default:
			fmt.Println("HandleRequest Running")
			time.Sleep(3 * time.Second)
		}
	}
}
