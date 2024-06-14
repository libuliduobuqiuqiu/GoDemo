package reflectdemo

import (
	"fmt"
	"reflect"
)

func BaseUseReflectDeepEqual() {

	a := make([]int, 100)
	b := make([]int, 100)
	fmt.Println(reflect.DeepEqual(a, b))

	a = append(a, 10)
	fmt.Println(reflect.DeepEqual(a, b))
}
