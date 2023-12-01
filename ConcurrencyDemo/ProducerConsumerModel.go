package ConcurrencyDemo

import (
	"fmt"
	"sync"
	"time"
)

// 生产者
func Producer(workerID int, ch chan<- int, producerWG *sync.WaitGroup) {
	defer producerWG.Done()

	for i := 0; i < 10; i++ {
		ch <- i * i
		fmt.Printf("Produce%d: %d\n", workerID, i*i)
		time.Sleep(2 * time.Second)
	}
}

// 消费者
func Consumer(workerID int, ch <-chan int, results chan<- int, consumerWG *sync.WaitGroup) {
	defer consumerWG.Done()

	for result := range ch {
		results <- result * 2
		fmt.Printf("Consumer %d get result: %d\n", workerID, result*2)
	}
}

func ExecModel() {
	consumerWG := &sync.WaitGroup{}
	producerWG := &sync.WaitGroup{}

	ch := make(chan int, 2)
	results := make(chan int)

	producerWG.Add(2)
	go Producer(1, ch, producerWG)
	go Producer(2, ch, producerWG)

	for i := 0; i < 5; i++ {
		consumerWG.Add(1)
		go Consumer(i, ch, results, consumerWG)
	}

	go func() {
		producerWG.Wait()
		close(ch)
	}()

	go func() {
		consumerWG.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Get Final Result: %d\n", result)
	}

}
