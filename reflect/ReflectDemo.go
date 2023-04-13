package main

import (
	"fmt"
	"reflect"
)

type Action interface {
	Hello()
}

type Person struct {
	Name string
	Age  int
}

type Animal struct {
	Kind string
	Name string
	Age  int
}

func (p *Person) Hello() {
	word := fmt.Sprintf("My name is %s, %d years old, I'm a Person", p.Name, p.Age)
	fmt.Println(word)
}

func (a *Animal) Hello() {
	word := fmt.Sprintf("My name is %s, %d years old, I'm a %s", a.Name, a.Age, a.Kind)
	fmt.Println(word)
}

func HandleReflect(r interface{}) (err error) {

	t := reflect.TypeOf(r)
	switch t.Kind() {
	case reflect.Ptr:

		v := reflect.ValueOf(r)
		temp := v.Elem()

		name := temp.FieldByName("Name")
		if name.IsValid() && name.CanSet() && name.Kind() == reflect.String {
			name.SetString(name.String() + "_flag")
		}

		hello := v.MethodByName("Hello")
		if hello.IsValid() {
			hello.Call(nil)
		}
		fmt.Println(v.String(), v.Interface().(*Person))
	default:
		fmt.Printf("不支持该%s类型对象操作", t.Kind())
	}

	return nil
}

func main() {
	// 创建一个Person类型的指针对象p，并赋值为&Person{"Alice", 18}
	p := &Person{"Alice", 18}
	if err := HandleReflect(p); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)

	a := &Animal{"Lion", "king", 10}
	if err := HandleReflect(a); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)

}
