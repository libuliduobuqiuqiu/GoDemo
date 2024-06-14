package reflectdemo

import (
	"fmt"
	"reflect"
)

func BaseUseReflectValue() {
	tmpstruct := &myStruct{Name: "linshukai"}

	rValue := reflect.ValueOf(tmpstruct)
	fmt.Println(rValue.Type())
	fmt.Println(rValue.Kind())

	s := rValue.Elem()
	fmt.Println(s)
	nameField := s.FieldByName("Name")
	nameField.SetString("zhangsan")
	fmt.Println(s.Interface())
}
