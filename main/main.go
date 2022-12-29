package main

import (
	"fmt"
	"sunrun/ADops"
)

func main() {
	a := `-\|/`
	for _, r := range `-\|/` {
		fmt.Printf("%c \n", r)
	}
	fmt.Printf("%T %s \n", a, a)

	ADops.Hello()
}
