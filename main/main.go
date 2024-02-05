package main

import (
	"fmt"
	"strings"
)

type linshukaiFloat float64

func getError() (string, error) {
	return "hello,world", fmt.Errorf("test error")
}

func getTmp() error {
	return fmt.Errorf("tmp error")
}

func get() (err error) {
	var err2 error
	fmt.Printf("%p", &err2)
	s, err2 := getError()
	if err2 != nil {
		fmt.Println(s)
		fmt.Printf("%p", &err2)
		fmt.Println(err2.Error())
		return err2
	}
	return
}

func main() {
	// httpdemo.HandleHttpRequest()
	a := linshukaiFloat(9.888)
	fmt.Println(a)
	b := float64(999.9)

	c := linshukaiFloat(b)
	fmt.Println(c)

	name := "Hello,world"
	strings.LastIndex(name, "Hello")
}
