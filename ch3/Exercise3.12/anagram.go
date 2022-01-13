package anagram

import (
	"reflect"
	"sort"
)

func IsAnagram(s1 string, s2 string) bool {
	slice1 := []rune(s1)
	slice2 := []rune(s2)
	sort.Slice(slice1, func(i int, j int) bool { return slice1[i] < slice1[j] })
	sort.Slice(slice2, func(i int, j int) bool { return slice2[i] < slice2[j] })

	return reflect.DeepEqual(slice1, slice2)
}
