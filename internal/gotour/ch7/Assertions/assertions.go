package main

import (
	"fmt"
)

type I interface {
	walk()
	speak()
}

type student struct {
	name    string
	age     string
	country string
}

type Teacher struct {
	name    string
	age     string
	country string
}

func (s student) walk() {
	fmt.Printf("I walk in my country: %s\n", s.country)
}

func (s student) speak() {
	fmt.Printf("My name is %s, I'm %s years old\n", s.name, s.age)
}

func (s Teacher) walk() {
	fmt.Printf("Teacher walk in my country: %s\n", s.country)
}

func (s Teacher) speak() {
	fmt.Printf("Teacher's name is %s, I'm %s years old\n", s.name, s.age)
}

func main() {
	var s I
	// 接口类型
	s = student{name: "zhangsan", age: "22", country: "China"}
	fmt.Printf("%s %T \n", s.(I), s.(I))

	s = Teacher{name: "wangwu", age: "42", country: "China Hongkong"}
	fmt.Printf("%s %T \n", s.(I), s.(I))

	// 非接口类型
	r, ok := s.(student)
	fmt.Println(r, ".", ok)

	var a I
	b := a.(I)
	fmt.Println(b)
}
