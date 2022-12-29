package main

import "fmt"

type DBTagInfo struct {
	InName  string
	OutName string
	Path    string
	In      bool
	Out     bool
}

func main() {
	params := []string{"zhangsan", "lisi", "wangwu"}
	list := []string{}
	list = append(list, params...)
	fmt.Println(list)
}
