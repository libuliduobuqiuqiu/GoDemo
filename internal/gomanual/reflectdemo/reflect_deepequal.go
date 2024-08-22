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

	var c, d int
	c = 100
	d = 100
	fmt.Printf("%p\n", &c)
	fmt.Printf("%p\n", &d)
	fmt.Println(reflect.DeepEqual(&c, &d))
}
