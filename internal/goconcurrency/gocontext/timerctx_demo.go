package gocontext

import (
	"context"
	"fmt"
	"time"
)

// WaitForCancellation returns the reason a context stopped.
func WaitForCancellation(ctx context.Context) error {
	<-ctx.Done()
	return ctx.Err()
}

func UseTimerCtx() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	fmt.Println("上下文取消：", WaitForCancellation(ctx))
}

func GenNum(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		var n int
		for {
			// Check cancellation before offering another value.
			select {
			case <-ctx.Done():
				return
			default:
			}
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}()

	return ch
}

func UseCancelGenNum() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for n := range GenNum(ctx) {
		fmt.Println(n)
		if n == 7 {
			cancel()
			break
		}
	}
}
