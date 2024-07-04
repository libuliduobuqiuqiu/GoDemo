package goalgorithm

import (
	"fmt"
	"godemo/internal/goalgorithm"
	"testing"
)

func TestUseBubbleSort(t *testing.T) {
	a := []int{92, 3, 1, 32, 33, 111, 23, 4, 555, 11, 22, 33, 1}
	b := goalgorithm.BubbleSort(a)
	fmt.Println(b)
}
