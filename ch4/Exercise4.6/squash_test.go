package squash

import (
	"bytes"
	"testing"
)

func TestSquash(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte("哈哈  哈 哈哈  a"), []byte("哈哈 哈 哈哈 a")},
		{[]byte(" 哈哈  哈 哈哈  a"), []byte(" 哈哈 哈 哈哈 a")},
		{[]byte("  哈哈  哈 哈哈  a"), []byte(" 哈哈 哈 哈哈 a")},
		{[]byte("哈哈  哈 哈哈  a "), []byte("哈哈 哈 哈哈 a ")},
		{[]byte("哈哈  哈 哈哈  a    "), []byte("哈哈 哈 哈哈 a ")},
	}

	for _, test := range tests {
		input := make([]byte, len(test.input))
		copy(input, test.input)
		got := Squash(test.input)
		if !bytes.Equal(got, test.want) {
			t.Errorf("Squash(%v) = %v, want=%v", input, got, test.want)
		}
	}
}
