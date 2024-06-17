package interview

import "testing"

// 测试回文函数是否正常显示

func IsPalindrom(s string) bool {
	for k := range s {
		if s[k] != s[len(s)-1-k] {
			return false
		}
	}
	return true
}

func TestPalindrom(t *testing.T) {
	a := "keyoa"

	if IsPalindrom(a) {
		t.Errorf("Error Function: %s", a)
	}

	b := "kakak"
	if IsPalindrom(b) {
		t.Errorf("Error Function: %s", b)
	}
}

func BenchmarkIsPalindrom(t *testing.B) {
	for i := 0; i < 10000; i++ {
		if !IsPalindrom("ababa") {
			t.Errorf("Error Function: %s", "ababa")
		}
	}
}

func TestCountLength(t *testing.T) {
	a := "zhangsan"

	if len(a) < 10 {
		t.Errorf("Error String length: %d", len(a))
	}
}
