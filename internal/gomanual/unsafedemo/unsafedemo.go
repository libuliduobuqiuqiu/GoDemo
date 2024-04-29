package unsafedemo

import (
	"fmt"
	"unsafe"
)

func UseUnsafePointer() {
	a := &map[string]string{}
	fmt.Println(a)
	b := unsafe.Pointer(a)
	fmt.Println(b)
}
