package methoddemo

import (
	"fmt"
	"reflect"
)

type ChinaPerson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c ChinaPerson) GrowUp() {
	c.Age += 1
}

func (c *ChinaPerson) GrowUpTrue() {
	c.Age += 1
}

func (c ChinaPerson) howOld() {
	fmt.Println(c.Name, ":", c.Age)
}

func (c ChinaPerson) HelloWorld() {
	fmt.Println("hello, " + c.Name)
}

func MethodUseDiffReceiver() {
	c := ChinaPerson{Name: "zhangsan", Age: 12}
	c.GrowUp()
	c.howOld()

	c.GrowUpTrue()
	c.howOld()

	c2 := &ChinaPerson{Name: "lisi", Age: 12}
	c2.GrowUp()
	c2.howOld()

	c2.GrowUpTrue()
	c2.howOld()

	t := reflect.TypeOf(c2)
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Index, t.Method(i).Name, t.Method(i).Type)
	}
}
