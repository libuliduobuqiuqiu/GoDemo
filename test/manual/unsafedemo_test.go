package manual

import (
	"godemo/internal/gomanual/unsafedemo"
	"testing"
)

func TestUnsafeDemo(t *testing.T) {

	unsafedemo.UseUnsafePointer()

}

func BenchmarkUnsafeDemo(b *testing.B) {

	var result []int
	for i := 0; i < 100000; i++ {
		x := i * i
		result = append(result, x)
	}

}
