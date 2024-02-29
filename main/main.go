package main

import (
	"fmt"
	"sunrun/gomanual/reflectdemo"
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
	// goconcurrency.StartChat()

	// goconcurrency.CountBalance()

	reflectdemo.ReflectPersonInfo(reflectdemo.HelloPersonInfo)
	node := Node{IP: "localhost"}
	P := reflectdemo.PersonInfo{Name: "linshukai", Age: 22}
	reflectdemo.ReflectPersonInfo(node)

	reflectdemo.SelectStructMethod(&P)
}
