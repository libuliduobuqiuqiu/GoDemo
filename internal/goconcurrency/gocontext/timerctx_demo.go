package gocontext

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func UseTimerCtx() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer cancel()
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("上下文取消：", ctx.Err())
				return
			default:
				fmt.Println("等待取消中。。。")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}

func GenNum(ctx context.Context) chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				close(ch)
				fmt.Println("Done.")
				return
			case ch <- n:
				n++
				time.Sleep(1 * time.Second)
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
	time.Sleep(10 * time.Second)
}
