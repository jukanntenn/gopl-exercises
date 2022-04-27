// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
	"math"

	popcount "github.com/jukanntenn/gopl-exercises/ch2/Exercise2.3"
)

func main() {
	var x1 IntSet
	x1.AddAll(1, 144, 9)
	fmt.Println(x1.String())
	var x2 IntSet
	x2.AddAll(2, 144)
	fmt.Println(x2.String())
	x1.IntersectWith(&x2)
	fmt.Println(x1.String())
	fmt.Println()

	var y1 IntSet
	y1.AddAll(1, 144, 9)
	fmt.Println(y1.String())
	var y2 IntSet
	y2.AddAll(2, 144)
	fmt.Println(y2.String())
	y1.DifferenceWith(&y2)
	fmt.Println(y1.String())
	fmt.Println()

	var z1 IntSet
	z1.AddAll(1, 144, 9)
	fmt.Println(z1.String())
	var z2 IntSet
	z2.AddAll(2, 144, 9, 13)
	fmt.Println(z2.String())
	z1.SymmetricDifference(&z2)
	fmt.Println(z1.String())
	fmt.Println()

	var s IntSet
	s.AddAll(1, 144, 9, 2, 13)
	for _, elem := range s.Elems() {
		println(elem)
	}
}

const size int = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/size, uint(x%size)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/size, uint(x%size)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", size*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	n := 0
	for _, word := range s.words {
		n += popcount.PopCount(uint64(word))
	}
	return n
}

func (s *IntSet) Remove(x int) {
	word, bit := x/size, uint(x%size)
	for word >= len(s.words) {
		return
	}
	op := uint(math.Pow(2, float64(bit)))
	s.words[word] &= (^op)
}

func (s *IntSet) Clear() {
	var op uint = 0
	for i := range s.words {
		s.words[i] &= op
	}
}

func (s *IntSet) Copy() *IntSet {
	n := len(s.words)
	words := make([]uint, n)
	for i, word := range s.words {
		words[i] = word
	}
	return &IntSet{words}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) Unionwith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Elems() []int {
	elems := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, size*i+j)
			}
		}
	}
	return elems
}
