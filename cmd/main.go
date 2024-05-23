package main

import "fmt"

type Node struct {
	FileNodeIP string
	Status     string
	NodeInfo   NodeInfo
}

type NodeInfo struct {
	ID      string
	IsVaild bool
}

func main() {
	n := Node{
		FileNodeIP: "10.21.21.97",
		Status:     "enabled",
	}
	fmt.Println(n.NodeInfo.ID)
	fmt.Println(n.NodeInfo.IsVaild)

	fmt.Println(100_100.)
}
