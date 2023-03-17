package reflect

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Hello() {
	p.Age += 1
}

func main() {
	// 创建一个Person类型的指针对象p，并赋值为&Person{"Alice", 18}
	p := &Person{"Alice", 18}

	t := reflect.TypeOf(*p)
	newP := reflect.New(t)
	fmt.Println("newP Type: ", newP.Kind(), newP.Kind() == reflect.Ptr)

	v := newP.Elem()
	fmt.Println("v Type: ", v.Kind(), v.Kind() == reflect.Struct)

	v2 := newP.Elem()
	fmt.Println("v2 Type ", v2.Kind(), v2.Kind() == reflect.Struct)

	// 打印v是否为指针类型
	//fmt.Println("v is pointer:", v.Kind() == reflect.Ptr) // v is pointer: true
	//// 获取v所指向值（即Person结构体）中名为name的字段f，并打印它所存储的字符串
	f := v2.FieldByName("Name")
	f.SetString("zhangsan")

	f2 := v2.FieldByName("Age")
	f2.SetInt(23)

	fmt.Println(v.Type())
	f3 := newP.MethodByName("Hello")
	f3.Call(nil)
	fmt.Println(p)
	fmt.Println(v)
}
