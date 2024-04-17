package main

import (
	"fmt"
	"sync"
)

// "github.com/beego/beego/v2/server/web"
// _ "godemo/internal/goweb/gowebsockets"
//
//

var numCh = make(chan struct{})
var strCh = make(chan struct{})

func printNum(wg *sync.WaitGroup, nums []int) {
	defer wg.Done()

	var total int
	for total = 0; total < len(nums); total++ {
		numCh <- struct{}{}
		fmt.Println(nums[total])

		if _, ok := <-strCh; !ok {
			fmt.Println("test2")
			break
		}
	}

	close(numCh)

	if total < len(nums) {
		for _, v := range nums[total:] {
			fmt.Println(v)
		}
	}

}

func printString(wg *sync.WaitGroup, strs []string) {
	defer wg.Done()

	var total int
	for total = 0; total < len(strs); total++ {
		if _, ok := <-numCh; !ok {
			fmt.Println("test")
			break
		}
		fmt.Println(strs[total])
		strCh <- struct{}{}
	}

	close(strCh)
	if total < len(strs) {
		for _, v := range strs[total:] {
			fmt.Println(v)
		}
	}
}

func main() {
	// web.BConfig.CopyRequestBody = true
	// web.Run(":8090")
	// goconcurrency.PrintFib()
	//
	wg := &sync.WaitGroup{}
	wg.Add(2)
	nums := []int{123123, 1232, 22, 64, 32, 2, 34, 65, 78, 1}
	strs := []string{"zhangsan", "linshu", "dsfaiosdjf", "dd", "3dsdf", "sdfaos"}
	go printNum(wg, nums)
	go printString(wg, strs)

	wg.Wait()
}
