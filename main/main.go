package main

import (
	"fmt"
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

func main() {

	ch := make(chan int)
	getNum2(ch)
	close(ch)
}
