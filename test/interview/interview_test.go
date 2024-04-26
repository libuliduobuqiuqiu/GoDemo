package interview_test

import (
	"fmt"
	"testing"
)

func ReverseStr(s string) string {

	tmp := []rune(s)
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	return string(tmp)
}

func TestReverseStr(t *testing.T) {
	var s = "linshukai"
	fmt.Println(ReverseStr(s))
}
