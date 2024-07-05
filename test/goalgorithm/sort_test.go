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

func TestUseInsertionSort(t *testing.T) {
	a := []int{92, 3, 1, 32, 33, 111, 23, 4, 555, 11, 22, 33, 1}
	b := goalgorithm.InsertionSort(a)
	fmt.Println(b)
}

func TestUseMergeSort(t *testing.T) {
	a := []int{92, 3, 1, 32, 33, 111, 23, 4, 555, 11, 22, 33, 1}
	b := goalgorithm.MergeSort(a)
	fmt.Println(b)
}

func TestUseQuickSort(t *testing.T) {
	a := []int{92, 3, 1, 32, 33, 111, 23, 4, 555, 11, 22, 33, 1}
	b := goalgorithm.QuickSort(a)
	fmt.Println(b)
}

func TestUseSelectionsSort(t *testing.T) {
	a := []int{92, 3, 1, 32, 33, 111, 23, 4, 555, 11, 22, 33, 1}
	b := goalgorithm.SelectionSort(a)
	fmt.Println(b)
}
