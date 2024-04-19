package gocontext

import (
	"context"
	"fmt"
	"time"
)

func genNum(ctx context.Context) <-chan int {
	var ch = make(chan int)
	var n int
	go func() {
		for {
			select {
			// 接受context的取消信号
			case <-ctx.Done():
				fmt.Println("goroutine exist")
				return
			case ch <- n:
				n++
				time.Sleep(time.Second)
			}
		}
	}()

	return ch
}

func RunGenNum() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for num := range genNum(ctx) {
		fmt.Println(num)
		if num > 4 {
			cancel()
			fmt.Println(ctx.Deadline())
			return
		}
	}
}
