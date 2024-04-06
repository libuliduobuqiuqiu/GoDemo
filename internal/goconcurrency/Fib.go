package goconcurrency

import (
	"fmt"
	"sync"
	"time"
)

var result = sync.Map{}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}

}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// 测试普通Map同时并发写入出现问题？go 自带-race机制检测数据竞争问题
func PrintFib() {
	var wg = &sync.WaitGroup{}
	go spinner(100 * time.Millisecond)
	execFib := func(n int, wg *sync.WaitGroup) {
		defer wg.Done()
		tmp := fib(n)
		fmt.Println(n, ":", tmp)
		result.Store(n, tmp)
	}

	var i int
	for i = 0; i < 50; i++ {
		wg.Add(1)
		go execFib(i, wg)
	}

	wg.Wait()
}
