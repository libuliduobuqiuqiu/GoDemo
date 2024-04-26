package reflectdemo

import (
	"fmt"
	"reflect"
)

type ObjOperation interface {
	GetOptions() map[string]string
}

type Man struct {
	sex string
	Person
}

type Person struct {
	Name    string
	Age     int
	Options map[string]string
}

func (p *Person) GetOptions() map[string]string {
	return p.Options
}

func (p *Person) SetOptions(o map[string]string) {
	p.Options = o
}
func ReflectUseMethod() {

	p := Person{Name: "zhangsan", Age: 22}
	p.SetOptions(map[string]string{"address": "guangdong"})

	m := Man{sex: "chengren", Person: p}
	val := reflect.ValueOf(m)

	field := val.FieldByName("Person")
	f := field.Interface()
	fmt.Printf("%+v, %T", f, f)

	if obj, ok := f.(ObjOperation); ok {
		fmt.Println(obj.GetOptions())
	}

	v := val.Interface()
	if obj, ok := v.(ObjOperation); ok {
		fmt.Println(obj.GetOptions())
	}
}

func ReflectChangeValue() {
	// 创建一个SPerson类型的指针对象p，并赋值为&SPerson{"Alice", 18}
	p := &SPerson{"Alice", 18}
	if err := HandleReflect(p); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)

	a := &Animal{"Lion", "king", 10}
	fmt.Println(reflect.TypeOf(a))
	if err := HandleReflect(a); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)

}
