package concurrency_test

import (
	"fmt"
	"godemo/internal/goconcurrency"
	"sort"
	"testing"
)

func TestUseContext(t *testing.T) {
	goconcurrency.UseContext()
}

func TestSortMap(t *testing.T) {
	a := map[string]int{
		"zhangsan":   11,
		"lisi":       12,
		"zhaozhenji": 33,
		"hairui":     33,
		"anhui":      34,
	}
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}
