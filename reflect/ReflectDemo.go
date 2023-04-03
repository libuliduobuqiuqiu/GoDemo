package reflect

import (
	"fmt"
	"reflect"
)

func SkipNumField(conf interface{}) {
	var v reflect.Value
	if r, ok := conf.(ObjOperation); ok {
		v = reflect.ValueOf(r).Elem()
	}

	t := reflect.TypeOf(conf)
	fmt.Println(t, t.Kind())




	fmt.Println(v, v.Type())

	for i:=0; i<v.NumField(); i++ {
		f := v.Field(i)
		fmt.Println(f.Type(), f.Interface())
	}

	n := v.FieldByName("Name")
	if n.IsValid(){
		fmt.Println(n.CanSet())
		fmt.Println(n.Type(), n.Interface())
		n.SetString("niubi")
	}



}