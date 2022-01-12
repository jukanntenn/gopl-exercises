package comma

import (
	"bytes"
)

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	buf.WriteByte(s[0])
	for i := 1; i < n; i++ {
		if (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
