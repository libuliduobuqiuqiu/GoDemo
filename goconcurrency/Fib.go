package goconcurrency

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {

	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}

}

func fib(n int) int {

	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func PrintFib() {
	go spinner(100 * time.Millisecond)
	fibN := fib(45)
	fmt.Printf("\rFibonacci (%d) = %d\n", 45, fibN)
}
