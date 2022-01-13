package anagram

import "testing"

func TestComma(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"a", "a", true},
		{"a", "b", false},
		{"abcde", "cbaed", true},
		{"abcde", "cbaez", false},
		{"abc", "bcad", false},
		{"abc", "abc", true},
	}

	for _, test := range tests {
		got := IsAnagram(test.s1, test.s2)
		if got != test.want {
			t.Errorf("IsAnagram(%s, %s) = %v", test.s1, test.s2, got)
		}
	}
}
