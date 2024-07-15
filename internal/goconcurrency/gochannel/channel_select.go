package gochannel

import (
	"fmt"
	"time"

	"github.com/go-faker/faker/v4"
)

func UseChannelSelect() {

	chInt := make(chan int)
	chString := make(chan string)

	fmt.Println("Start.")

	select {
	case i := <-chInt:
		fmt.Println(i)
	case j := <-chString:
		fmt.Println(j)
	case <-time.After(5 * time.Second):
		fmt.Println("select exit.")
	}

	fmt.Println("Done.")
}

func UseChannelDone() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	for {
		select {
		case <-ch:
			fmt.Println("Done.")
			return
		case <-time.After(10 * time.Second):
			fmt.Println("Timeout.")
			return
		}
	}
}

func UseChannelRange() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			data := faker.ChineseName()
			ch <- data
			time.Sleep(1 * time.Second)
		}
		fmt.Println("ch is closed")
		close(ch)
	}()

	for data := range ch {
		fmt.Println(data)
	}

	fmt.Println("Done")
}
