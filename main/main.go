package main

import (
	"fmt"
	// "sunrun/gostorage/standardmysql"
	"sync"
	"time"
)

func getNum1(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()

	for {
		select {
		case num := <-ch:
			fmt.Println("getNum1", num)
			return
		case <-time.After(10 * time.Millisecond):
			fmt.Println("TimeOut.")
		}
	}

}

func getNum2(ch chan int) {
	ch <- 3
	num := <-ch
	fmt.Println("getNum2", num)
}

func setNum(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	time.Sleep(20 * time.Millisecond)
	ch <- 2
}

var (
	a = make(chan struct{}, 1)
)

func addA(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	a <- struct{}{}
	fmt.Println(num)
	time.Sleep(2 * time.Second)
	<-a
}

type Node struct {
	IP string `json:"node"`
}

func main() {
	// 启动聊天室
	// goconcurrency.StartChat()

	// goconcurrency.CountBalance()

	// 测试使用内置sql引擎执行mysql语句
	// standardmysql.ExecSQLStr()
Out:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > 3 {
				continue Out
			}
			fmt.Println(i, j)
		}
	}
}
