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
