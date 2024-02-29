package reflectdemo

import (
	"fmt"
	"reflect"
)

type PersonInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *PersonInfo) PrintName() {
	fmt.Println(p.Name)
}

func (p *PersonInfo) PrintAge() {
	fmt.Println(p.Age)
}

func HelloPersonInfo(p *PersonInfo) {
	fmt.Println("Hllo, world ", p.Name)
}

// 反射函数执行
func ReflectPersonInfo(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println(t)
	if t.Kind() == reflect.Func {
		f := reflect.ValueOf(i)
		fmt.Println(f, f.Type())
		p := &PersonInfo{Name: "linshukai", Age: 12}
		params := []reflect.Value{}
		params = append(params, reflect.ValueOf(p))
		f.Call(params)
	} else {
		fmt.Println(t.Name())
	}
}

func SelectStructMethod(i interface{}) {
	tmp := reflect.ValueOf(i)
	tmpType := tmp.Type()
	fmt.Println(tmp.Kind())
	if tmp.Elem().Kind() == reflect.Struct {
		for k := 0; k < tmp.Elem().NumField(); k++ {
			f := tmp.Method(k)
			fmt.Println(tmpType, tmpType.Method(k).Name, f.Type())
		}
	}
}
