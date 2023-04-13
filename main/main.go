package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	s := "how old are you"
	rand.Seed(time.Now().UnixNano())
	for k, v := range s {
		fmt.Println(k, v)
		fmt.Println(rand.Intn(int(v)))
	}
}
