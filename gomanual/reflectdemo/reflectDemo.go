package reflectdemo

import (
	"fmt"
	"reflect"
)

type Action interface {
	Hello()
}

type SPerson struct {
	Name string
	Age  int
}

type Animal struct {
	Kind string
	Name string
	Age  int
}

func (p *SPerson) Hello() {
	word := fmt.Sprintf("My name is %s, %d years old, I'm a SPerson", p.Name, p.Age)
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
		fmt.Println(v.String(), v.Interface().(*SPerson))
	default:
		fmt.Printf("不支持该%s类型对象操作", t.Kind())
	}

	return nil
}
