package test

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

type MainPerson struct {
	Name string `json:"name"`
}

func JudgeMainPerson(p interface{}) {
	if person, ok := p.(*MainPerson); ok {
		fmt.Println(person.Name)
		person.Name = "linsan"
	}
}

func JudgeList(notExistPerson []string) {
	fmt.Println(len(notExistPerson))
	if len(notExistPerson) > 0 {
		fmt.Println("Not Exist Person")
	}
	fmt.Println("done.")
}

func TestJudegeMainPerson(t *testing.T) {
	p := MainPerson{
		Name: "zhangsan",
	}
	JudgeMainPerson(&p)
	fmt.Println(p.Name)
}

func TestJudgeList(t *testing.T) {
	JudgeList(nil)
}

func ReturnStruct() (data MainPerson) {
	fmt.Printf("%p\n", &data)
	return
}

func TestReturnStruct(t *testing.T) {
	data := ReturnStruct()
	fmt.Printf("%p\n", &data)
	return
}

func TestSliceAppend(t *testing.T) {
	var b []int = nil
	fmt.Printf("%p\n", b)
	fmt.Println(len(b), cap(b))

	b = append(b, 100)
	b = append(b, 200, 200, 200, 200, 200, 200)
	fmt.Printf("%p\n", b)
	fmt.Println(len(b), cap(b))

	s := make([]int, 2, 2)
	s[0] = 1
	s[1] = 2

	// 打印初始切片的地址和底层数组的地址
	fmt.Printf("Initial slice address: %p\n", s)
	fmt.Printf("Initial array address: %p\n", unsafe.Pointer(&s[0]))

	// 添加元素，触发扩容
	s = append(s, 3)

	// 打印扩容后的切片地址和底层数组的地址
	fmt.Printf("New slice address: %p\n", s)
	fmt.Printf("New array address: %p\n", unsafe.Pointer(&s[0]))
}

func ChangeSlice(a []int) []int {
	a[1] = 99999
	a = append(a, 100000)
	fmt.Println(len(a), cap(a))
	return a
}

func TestSliceSend(t *testing.T) {

	a := []int{10, 10}
	a = append(a, 10)
	b := ChangeSlice(a)
	fmt.Println(a, b)
	fmt.Println(len(a), cap(a))
}

func TestRange(t *testing.T) {

	values := []string{"zhangsan", "wangwu"}
	for _, i := range values[3:] {
		fmt.Println(i)
	}
}

type person struct {
	Name string
	Age  int
}

func TestSliceChange(t *testing.T) {
	pList := []person{
		{Name: "t1", Age: 22},
	}

	for _, p := range pList {
		p.Age = ChangePerson(p.Age)
	}

	fmt.Println(pList)

}

func ChangePerson(p int) int {
	return p + 1
}

func TestForPrint(t *testing.T) {
	a := 100
	go func() {
		time.Sleep(2 * time.Second)
		a = 200
	}()

	for a != 200 {
		fmt.Println("yes,ok")
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("no,i'm right")
}
