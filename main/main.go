package main

import (
	"fmt"
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

func reverseList(l [10]*int) {
	for i, j := 0, len(l)-1; i <= j; i, j = i+1, j-1 {
		*l[i], *l[j] = *l[j], *l[i]
	}
}

func reverseSlice(l []int) {
	for i, j := 0, len(l)-1; i <= j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
}

func changeList(data []string) {
	data[0] = "zhaoyun"
	fmt.Println(cap(data), len(data))
	data = append(data, "zhangsan")
	data[0] = "zhuangzi"
}

func main() {
	data := []string{"wangwu", "lisi"}
	changeList(data)
	fmt.Println(data)
	// httpdemo.RequestHtml("https://books.studygolang.com/gopl-zh/ch5/ch5-02.html")
}
