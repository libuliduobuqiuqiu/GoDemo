package manual

import "testing"

func Reverse(s string) string {
	tmp := []rune(s)
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return string(tmp)
}

func FuzzReverse(f *testing.F) {
	data := []string{
		"hello,world",
		"what",
		"fuzzy",
	}

	for _, v := range data {
		f.Add(v)
	}

	f.Fuzz(func(t *testing.T, a string) {

		second := Reverse(a)
		third := Reverse(second)
		t.Log(third)
	})
}
