package test

import (
	"fmt"
	"godemo/internal/godemo"
	"testing"
)

func TestBytesConv(t *testing.T) {

	godemo.CountParams("/web/user/info/:id")

}

func TestAddSlice(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}
	var l [5]int

	// 切片本质底层是指针，range的时候复制的是指针，操作切片则反馈到底层数组，而数组赋值的时候
	// 也会获取到底层的数组的值
	for k, v := range s {
		if k == 0 {
			s[1] = 12
			s[2] = 13
		}
		l[k] = v
	}

	fmt.Println(s)
	fmt.Println(l)
}
