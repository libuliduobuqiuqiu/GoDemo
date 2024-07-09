package gochannel

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-faker/faker/v4"
)

var closedCh chan struct{}

func producer(producerId int, ch chan<- string) {

	for {
		data := faker.ChineseName()

		select {
		case ch <- data:
			fmt.Printf("Producer %d : %s \n", producerId, data)
		case <-closedCh:
			fmt.Printf("Producer %d exist\n", producerId)
			return
		}
	}
}

func consumer(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("Consumer Get %s \n", data)
			if count == 10 {
				close(closedCh)
				time.Sleep(1 * time.Second)
				return
			}
			count += 1
			time.Sleep(1 * time.Second)
		}
	}
}

func UseProducerConsumer() {
	wg := &sync.WaitGroup{}
	closedCh = make(chan struct{})
	ch := make(chan string, 3)
	for i := 0; i < 20; i++ {
		go producer(i, ch)
	}

	wg.Add(1)
	go consumer(ch, wg)
	wg.Wait()
}
