package reverse

import (
	"bytes"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte("一二三"), []byte("三二一")},
		{[]byte("你好，世界"), []byte("界世，好你")},
		{[]byte(" 你 好 ，世界"), []byte("界世， 好 你 ")},
	}

	for _, test := range tests {
		input := make([]byte, len(test.input))
		copy(input, test.input)
		got := ReverseUTF8(test.input)
		if !bytes.Equal(got, test.want) {
			t.Errorf("Reverse(%v) = %v, want=%v", input, got, test.want)
		}
	}
}
