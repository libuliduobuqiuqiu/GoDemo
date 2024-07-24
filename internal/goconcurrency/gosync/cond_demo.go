package gosync

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
)

var status int64

// 通过条件变量sync.Cond,Wait()释放锁后等待唤醒，Broadcast()广播唤醒。
// 启用多个监听，等待广播完毕后Goroutine继续执行
func UseSyncCond() {
	c := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 10; i++ {
		go Listener(c)
	}

	go Broadcaster(c)

	// 监听ctrl-c输入停止
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func Listener(c *sync.Cond) {
	c.L.Lock()

	for atomic.LoadInt64(&status) != 1 {
		c.Wait()
	}

	fmt.Println("Listening")
	c.L.Unlock()
}

func Broadcaster(c *sync.Cond) {
	c.L.Lock()

	atomic.StoreInt64(&status, 1)
	c.Broadcast()
	c.L.Unlock()
}
