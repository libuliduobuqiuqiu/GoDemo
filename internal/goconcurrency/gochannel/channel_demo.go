package gochannel

import (
	"fmt"
	"math/rand"
	"time"
)

type ChUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var pu = &ChUser{Name: "linshukai", Age: 22}
var g chan *ChUser

// 测试通道在关闭之后还有数据，是否能够读出来
func RecvClosedChannel() {
	ch := make(chan string, 5)
	ch <- "hello"
	ch <- "world"
	ch <- "marry"
	text, ok := <-ch
	fmt.Println(text, ok)
	close(ch)
	for t := range ch {
		fmt.Println(t)
	}

	n, ok := <-ch
	fmt.Println(n, ok)
	fmt.Println(n == "")
	fmt.Println(rand.Intn(1000))
}

func PrintUser(g <-chan *ChUser) {
	user := <-g
	fmt.Printf("%p %v\n", user, *user)
}

func ModifyUser(u *ChUser) {
	fmt.Println("Modify Received From: ", u)
	u.Age = 200
}

// 测试通道发送和接收的本质：“值”的拷贝
func UseChannelPrintUser() {
	g = make(chan *ChUser, 2)
	g <- pu
	fmt.Printf("%p %v\n", pu, *pu)
	pu = &ChUser{Name: "zhangsan", Age: 22}
	go PrintUser(g)
	go ModifyUser(pu)
	time.Sleep(1 * time.Second)
	fmt.Println(pu)
}

// 定时器
func UseChannelTicker() {
	ticker := time.Tick(10 * time.Second)
	select {
	case <-ticker:
		fmt.Println("Timeout.")
		return
	}
}

func UseChannelTimer() {
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("10 Seconds Timeout.")
		return
	}
}
