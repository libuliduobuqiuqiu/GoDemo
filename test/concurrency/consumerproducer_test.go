package concurrency

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

var (
	data     chan interface{}
	producer *Producer
	consumer *Consumer
)

type Producer struct{}
type Consumer struct{}
type Message struct {
	Info string `json:"info"`
}

func (p *Producer) Produce(msg string, wg *sync.WaitGroup, data chan<- interface{}) {
	defer wg.Done()
	data <- Message{Info: msg}
}

func (c *Consumer) Consume(wg *sync.WaitGroup, data <-chan interface{}) {
	defer wg.Done()
	Msg := <-data

	msgJson, _ := json.Marshal(Msg)
	fmt.Println(string(msgJson))
}

func TestChan(t *testing.T) {

	wg := &sync.WaitGroup{}
	data := make(chan interface{})

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer.Produce(fmt.Sprintf("%d", i), wg, data)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go consumer.Consume(wg, data)
	}

}
