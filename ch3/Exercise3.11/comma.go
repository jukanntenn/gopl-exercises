package comma

import (
	"bytes"
	"fmt"
	"strings"

	commaint "github.com/jukanntenn/gopl-exercises/ch3/Exercise3.10"
)

func Comma(s string) string {
	if s == "" {
		return s
	}
	n := len(s)
	var buf bytes.Buffer
	sign := s[0]
	start := 0
	if sign == '+' || sign == '-' {
		buf.WriteByte(sign)
		start = 1
	}

	index := strings.Index(s, ".")
	end := n
	if index != -1 {
		end = index
	}
	fmt.Println(s[start:end])
	buf.WriteString(commaint.Comma(s[start:end]))
	buf.WriteString(s[end:])
	return buf.String()
}
