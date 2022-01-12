package comma

import "testing"

func TestComma(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1", "1"},
		{"814", "814"},
		{"5678", "5,678"},
		{"567891", "567,891"},
		{"99999999999999", "99,999,999,999,999"},
	}

	for _, test := range tests {
		got := Comma(test.input)
		if got != test.want {
			t.Errorf("Comma(%s) = %s", test.input, got)
		}
	}
}
