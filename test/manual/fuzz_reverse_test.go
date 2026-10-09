package manual

import (
	"testing"
	"unicode/utf8"
)

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
		if !utf8.ValidString(a) {
			t.Skip("Reverse operates on UTF-8 text")
		}
		second := Reverse(a)
		third := Reverse(second)
		if third != a {
			t.Fatalf("reverse twice: got %q, want %q", third, a)
		}
	})
}
