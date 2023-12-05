package main

import "fmt"

func main() {
	if err := InitDB(); err != nil {
		fmt.Println(err)
		return
	}
	QueryMultiRowDemo()
}
