package gochannel

import (
	"fmt"
	"time"
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
