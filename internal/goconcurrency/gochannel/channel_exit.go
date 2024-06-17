package gochannel

import (
	"fmt"
	"sync"
	"time"
)

func InputChannel(ch chan<- string) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("number: %d", i)
	}

	time.Sleep(5 * time.Second)
}

func OutputChannel(wg *sync.WaitGroup, ch <-chan string) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("Input Channel is exist.")
}

func UseChannel() {
	wg := &sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(1)
	go InputChannel(ch)
	go OutputChannel(wg, ch)

	wg.Wait()
}
