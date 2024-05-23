package manual

import (
	"godemo/internal/gomanual/genericsdemo"
	"testing"
)

func FuzzEqual(f *testing.F) {
	var data = [][]int{
		{1, 2},
		{2, 3},
		{4, 5},
	}
	for _, v := range data {
		f.Add(v[0], v[1])
	}

	f.Fuzz(func(t *testing.T, a, b int) {
		var flag1, flag2 bool
		if a == b {
			flag1 = true
		}

		flag2 = genericsdemo.Equal(a, b)

		if flag1 != flag2 {
			t.Fatalf("Equal(%d, %d) error", a, b)
		}
	})
}
