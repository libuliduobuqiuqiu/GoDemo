package reflectdemo

import (
	"fmt"
	"reflect"
)

type rFace interface {
	hello()
}

type myStruct struct {
	Name string `json:"Name,omiempty"`
}

func (m myStruct) hello() {
	fmt.Println("hello,world")
}

func BaseUseReflectType() {
	tmpMap := map[string]int{}
	rType := reflect.TypeOf(tmpMap)
	fmt.Println(rType.Kind())
	fmt.Println(rType.Elem())
	fmt.Println(rType.Size())

	tmpStruct := myStruct{}
	sType := reflect.TypeOf(tmpStruct)

	var tmpFace = new(rFace)
	rIface := reflect.TypeOf(tmpFace).Elem()
	fmt.Println(sType.Comparable())
	fmt.Println(sType.Implements(rIface))
}
