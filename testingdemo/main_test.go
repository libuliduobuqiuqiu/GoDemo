package testingdemo

import (
	"fmt"
	"testing"
)

func Add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("xxxx")
	}
}

func ExampleAdd() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(2, 2))

	// Output:
	// 3
	// 5
}
