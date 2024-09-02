package gochannel

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/go-faker/faker/v4"
)

var total int64

// 中间人角色
func mediator(stopedCh, toStopCh chan struct{}) {
	<-toStopCh
	close(stopedCh)
	return
}

// 发送者，统计发送两百条后停止发送
func sender(id int, ch chan string, toStopCh, stopedCh chan struct{}) {
	for {
		data := faker.ChineseName()

		// 统计发送200次之后停止
		if total == 200 {
			toStopCh <- struct{}{}
			return
		}

		select {
		case ch <- data:
			fmt.Printf("Sender %d send: %s successfully.\n", id, data)
			atomic.AddInt64(&total, 1)
		case <-stopedCh:
			fmt.Printf("Sender %d closed sucessfully.\n", id)
			return
		}
	}
}

// 接收者
func receiver(id int, ch chan string, stopedCh chan struct{}) {
	for {
		select {
		case data := <-ch:
			fmt.Printf("Receiver %d receive: %s successfully.\n", id, data)
			time.Sleep(1 * time.Second)
		case <-stopedCh:
			fmt.Printf("Receiver %d closed successfully.\n", id)
			return
		}
	}
}

func UseChannelClosedGracefully() {
	toStopCh := make(chan struct{})
	stopedCh := make(chan struct{})
	ch := make(chan string, 40)

	go mediator(stopedCh, toStopCh)
	for i := 0; i <= 20; i++ {
		go receiver(i, ch, stopedCh)
	}

	for i := 0; i <= 30; i++ {
		go sender(i, ch, toStopCh, stopedCh)
	}

	select {
	case <-time.After(60 * time.Second):
	}
}
