package eliminate

import (
	"testing"
)

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func TestEliminateAdjacentDups(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{[]string{}, []string{}},
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{[]string{"a", "a", "a", "a"}, []string{"a"}},
		{[]string{"a", "a", "b", "b"}, []string{"a", "b"}},
		{[]string{"a", "a", "b", "c", "c", "a", "b"}, []string{"a", "b", "c", "a", "b"}},
		{[]string{"a", "a", "b", "c", "d", "d", "d", "a", "b"}, []string{"a", "b", "c", "d", "a", "b"}},
	}

	for _, test := range tests {
		input := make([]string, len(test.input))
		copy(input, test.input)
		got := EliminateAdjacentDups(test.input)
		if !equal(got, test.want) {
			t.Errorf("TestEliminateAdjacentDups(%v) = %v, want=%v", input, got, test.want)
		}
	}
}
