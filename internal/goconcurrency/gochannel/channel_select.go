package gochannel

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-faker/faker/v4"
)

// SelectFirst waits for one value or cancellation. A nil channel disables its case.
func SelectFirst(ctx context.Context, numbers <-chan int, texts <-chan string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case number, ok := <-numbers:
		if !ok {
			return "", errors.New("number channel closed")
		}
		return strconv.Itoa(number), nil
	case value, ok := <-texts:
		if !ok {
			return "", errors.New("text channel closed")
		}
		return value, nil
	}
}

func UseChannelSelect() {
	numbers := make(chan int, 1)
	numbers <- 42
	value, err := SelectFirst(context.Background(), numbers, nil)
	fmt.Println(value, err)
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
