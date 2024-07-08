package gochannel

import (
	"fmt"
	"math/rand"
)

func RecvClosedChannel() {
	ch := make(chan string, 5)

	ch <- "hello"

	ch <- "world"

	ch <- "marry"

	text, ok := <-ch
	fmt.Println(text, ok)

	close(ch)

	for t := range ch {
		fmt.Println(t)
	}

	n, ok := <-ch
	fmt.Println(n, ok)
	fmt.Println(n == "")
	fmt.Println(rand.Intn(1000))
}
