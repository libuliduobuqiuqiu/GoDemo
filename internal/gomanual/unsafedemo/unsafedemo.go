package unsafedemo

import (
	"fmt"
	"unsafe"
)

type UnsafePerson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UnsafeChildren struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func UseUnsafeConvert() {
	p := UnsafePerson{
		Name: "zhangsan",
		Age:  22,
	}

	tmpCh := (*UnsafeChildren)(unsafe.Pointer(&p))
	fmt.Printf("%+v", tmpCh)
}
