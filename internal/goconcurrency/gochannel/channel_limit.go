package gochannel

import (
	"fmt"
	"sync"
	"time"
)

func PrintTask(task string, ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- struct{}{}
	time.Sleep(2 * time.Second)
	fmt.Println("Finished Task:", task)
	<-ch
}

// 通过通道控制并发数
func UseLimitGoroutine(limit int, tasks []string) {
	// wg用于控制并发，主协程等待所有完成后才退出
	wg := &sync.WaitGroup{}
	if limit == 0 {
		limit = 5
	}

	ch := make(chan struct{}, limit)
	for _, task := range tasks {
		wg.Add(1)
		go PrintTask(task, ch, wg)
	}

	wg.Wait()
}
