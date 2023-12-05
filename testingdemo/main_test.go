package testingdemo

import (
	"fmt"
	"testing"
)

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.Fatal("xxxx")
	}
}

func ExampleAdd() {
	fmt.Println(add(1, 2))
	fmt.Println(add(2, 2))

	// Output:
	// 3
	// 4
}
