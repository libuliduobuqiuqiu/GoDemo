package main

import (
	"flag"
	"fmt"
	"time"
)

var Period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v ....", *Period)
	time.Sleep(*Period)
	fmt.Println()
}
