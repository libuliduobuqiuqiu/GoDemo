package gosync

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

var status int64

// 通过条件变量sync.Cond,Wait()释放锁后等待唤醒，Broadcast()广播唤醒。
// 启用多个监听，等待广播完毕后Goroutine继续执行
func UseSyncCond() {
	c := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 10; i++ {
		go Listener(i, c)
	}

	go Broadcaster(c)

	// 监听ctrl-c输入停止
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func Listener(num int, c *sync.Cond) {
	c.L.Lock()
	fmt.Println("Lock", num)
	for atomic.LoadInt64(&status) != 1 {

		fmt.Println("Listening: ", num)
		c.Wait()
	}
	fmt.Println("Unlock", num)
	c.L.Unlock()
}

func Broadcaster(c *sync.Cond) {
	c.L.Lock()
	time.Sleep(2 * time.Second)
	atomic.StoreInt64(&status, 1)
	c.Broadcast()
	c.L.Unlock()
}
