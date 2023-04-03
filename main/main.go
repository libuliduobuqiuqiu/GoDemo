package main

import (
	"fmt"
	"sunrun/reflect"
)

type ObjOperation interface{
	GetOptions()map[string]string
}

type Man struct {
	sex string
	Person
}

type Person struct {
	Name string
	Age int
	Options map[string]string
}

func (p *Person) GetOptions() map[string]string {
	return p.Options
}

func (p *Person) SetOptions(o map[string]string){
	p.Options = o
}

func main (){
	info := map[string]string{"zhangsan": "wangwu", "lisi":"zhaosi"}
	p := &Person{"xjp", 12, info}
	reflect.SkipNumField(p)

	fmt.Println(p)
}