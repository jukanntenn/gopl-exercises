package main

import (
	"fmt"
	"sort"
)

func main() {
	s := str("abcdefgfedcba")
	fmt.Println(IsPalindrome(s))

	s = str("abccda")
	fmt.Println(IsPalindrome(s))

	s = str("你好世界世好你")
	fmt.Println(IsPalindrome(s))

	s = str("你好，世界！")
	fmt.Println(IsPalindrome(s))

	s = str("hello, 世界！！界世 ,olleh")
	fmt.Println(IsPalindrome(s))
}

type str []rune

func (s str) Len() int           { return len(s) }
func (s str) Less(i, j int) bool { return s[i] < s[j] }
func (s str) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
