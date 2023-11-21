package exercises

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if IsPalindrome("abc") {
		t.Fatalf(`Value("") = %v, want "%v"`, true, false)
	}

	if !IsPalindrome("abcba") {
		t.Fatalf(`Value("") = %v, want "%v"`, false, true)
	}

	if !IsPalindrome("日月光月日") {
		t.Fatalf(`Value("") = %v, want "%v"`, false, true)
	}
}
