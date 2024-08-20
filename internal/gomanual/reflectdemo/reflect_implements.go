package reflectdemo

import (
	"fmt"
	"reflect"
)

type Factory interface {
	Create()
	Count() int
}

type Aodi struct {
	Name string
	Cars int
}

func (a *Aodi) Create() {
	fmt.Println("Generate a new car.")
	a.Cars += 1
}

func (a *Aodi) Count() int {
	fmt.Println("Now have cars: ", a.Cars)
	return a.Cars
}

func SimpleFactory(f interface{}) {

	t := reflect.TypeOf(f)
	if t.Implements(reflect.TypeOf((*Factory)(nil)).Elem()) {
		fmt.Println("true")
		return
	}

	fmt.Println("false")
}

func ReflectImplments() {

	a := &Aodi{
		Name: "aodi",
	}

	SimpleFactory(a)

}
