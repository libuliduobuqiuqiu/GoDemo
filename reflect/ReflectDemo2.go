package reflect

import (
	"fmt"
	"reflect"
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


	p := Person{Name:"zhangsan", Age: 22}
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