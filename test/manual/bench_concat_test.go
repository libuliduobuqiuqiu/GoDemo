package manual

import (
	"strings"
	"testing"
)

var longString = "longStringlongStringlongStringlongStringlongStringlongStringlongStringlongString"

func ConcatStringDirct(s string) {
	var res string
	for i := 0; i < 100_00.; i++ {
		res += s
	}
}

func ConcatStringWithBuilder(s string) {
	var res strings.Builder

	for i := 0; i < 100_00.; i++ {
		res.WriteString(s)
	}
}

func BenchmarkConcatDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringDirct(longString)
	}
}

func BenchmarkConcatWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringWithBuilder(longString)
	}
}
