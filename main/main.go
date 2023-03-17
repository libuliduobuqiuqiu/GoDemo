package main

import (
	"fmt"
)

type info interface {
	SayName() string
	SayAge() int
}

type handle interface {
	HandleInfo()
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayName() string {
	return p.Name
}

func (p *Person) SayAge() int {
	return p.Age
}

func (p *Person) SetName(n string) {
	p.Name = n
}

func (p *Person) HandleInfo() {
	fmt.Println("My Name is ", p.Name, "My Age is  ", p.Age)
}

func Hello() *Person {
	fmt.Println("Hello,world")
	return &Person{"zhangsna", 22}
}

func main() {
	p2 := &Person{}
	p2.SetName("lisi")
	fmt.Println(p2.SayName())

	a := map[string]func() interface{}{
		"zhangsan": func() interface{} { return Hello() },
	}

	fmt.Println(a)
	fmt.Println(a["zhangsan"]())

}
