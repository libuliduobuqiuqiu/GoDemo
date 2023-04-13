package main

import (
	"fmt"
	"reflect"
)

type MyInterface interface {
	Method1()
	Method2()
}

type MyStruct struct {
}

func (s *MyStruct) Method1() {
	fmt.Println("Method1")
}

func (s *MyStruct) Method2() {
	fmt.Println("Method2")
}

func tmain() {
	var s MyStruct
	var i MyInterface = &s
	fmt.Println(reflect.TypeOf(i).Implements(reflect.TypeOf((*MyInterface)(nil)).Elem()))
}
