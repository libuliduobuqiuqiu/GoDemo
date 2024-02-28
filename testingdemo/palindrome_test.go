package testingdemo

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
