package main

import (
	"fmt"
	"log"
	"sunrun/gomanual/reflectdemo"
	"sunrun/gostorage/standardmysql"
	"sunrun/public"

	// "sunrun/gostorage/standardmysql"
	"sync"
	"time"
)

func getNum1(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()

	for {
		select {
		case num := <-ch:
			fmt.Println("getNum1", num)
			return
		case <-time.After(10 * time.Millisecond):
			fmt.Println("TimeOut.")
		}
	}

}

func getNum2(ch chan int) {
	ch <- 3
	num := <-ch
	fmt.Println("getNum2", num)
}

func setNum(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	time.Sleep(20 * time.Millisecond)
	ch <- 2
}

var (
	a = make(chan struct{}, 1)
)

func addA(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	a <- struct{}{}
	fmt.Println(num)
	time.Sleep(2 * time.Second)
	<-a
}

type Node struct {
	IP string `json:"node"`
}

func main() {
	// 启动聊天室
	// goconcurrency.StartChat()

	// goconcurrency.CountBalance()

	// 测试使用内置sql引擎执行mysql语句
	// standardmysql.ExecSQLStr()

	d := reflectdemo.Device{Name: "zhangsan", Address: "127.0.0.1", Port: 8080}
	sql, sqlValues, err := reflectdemo.GenerateMysqlString(&d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sql)

	globalConfig := public.GetGlobalConfig()
	db, err := standardmysql.GetDB(globalConfig.MysqlConfig)

	_, err = db.Exec(sql, sqlValues...)
	if err != nil {
		log.Fatal(err)
	}
}
