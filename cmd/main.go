package main

import (
	"fmt"
	"godemo/internal/godemo"
	"godemo/internal/gomanual/genericsdemo"
)

func main() {
	type a int8
	genericsdemo.PrintMan()

	err := godemo.HandleError()
	fmt.Println(err)
}
