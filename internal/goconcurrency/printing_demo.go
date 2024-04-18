package goconcurrency

import (
	"fmt"
	"sync"
)

var numCh = make(chan struct{})
var strCh = make(chan struct{})

// printLastData: 切片中剩余元素全部打印
func printLastData[T any](total int, data []T) {
	for ; total < len(data); total++ {
		fmt.Println(data[total])
	}
}

func printNum(wg *sync.WaitGroup, nums []int) {
	defer wg.Done()

	var total int
	// 先启动打印string的通道
	strCh <- struct{}{}

exit:
	for {
		select {

		case _, ok := <-numCh:
			if !ok {
				break exit
			}
			if total >= len(nums) {
				break exit
			}
			fmt.Println(nums[total])
			total++
			strCh <- struct{}{}
		}
	}

	close(strCh)
	printLastData[int](total, nums)
}

func printString(wg *sync.WaitGroup, strs []string) {
	defer wg.Done()
	var total int

exit:
	for {
		select {
		case _, ok := <-strCh:
			if !ok {
				break exit
			}

			if total >= len(strs) {
				break exit
			}
			fmt.Println(strs[total])
			total++
			numCh <- struct{}{}
		}
	}

	close(numCh)
	printLastData[string](total, strs)
}

func PrintingStrsNums() {
	// 交替打印数字和字符串
	wg := &sync.WaitGroup{}
	wg.Add(2)
	nums := []int{123123, 1232, 22, 64, 32, 2, 34, 65, 78, 1}
	strs := []string{"zhangsan", "linshu", "dsfaiosdjf", "dd", "3dsdf", "sdfaos"}
	go printNum(wg, nums)
	go printString(wg, strs)

	wg.Wait()
}
