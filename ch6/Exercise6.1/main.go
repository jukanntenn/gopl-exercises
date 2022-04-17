// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
	"math"

	popcount "github.com/jukanntenn/gopl-exercises/ch2/Exercise2.3"
)

func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	fmt.Printf("length of x: %d\n", x.Len())

	y := x.Copy()

	x.Remove(9)
	fmt.Println(x.String())
	fmt.Printf("length of x: %d\n", x.Len())
	fmt.Println(y.String())
	fmt.Printf("length of y: %d\n", y.Len())

	x.Clear()
	fmt.Println(x.String())
	fmt.Printf("length of x: %d\n", x.Len())

	y.Remove(144)
	fmt.Println(y.String())
	fmt.Printf("length of y: %d\n", y.Len())
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
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

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	n := 0
	for _, word := range s.words {
		n += popcount.PopCount(word)
	}
	return n
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		return
	}
	op := uint64(math.Pow(2, float64(bit)))
	s.words[word] &= (^op)
}

func (s *IntSet) Clear() {
	var op uint64 = 0
	for i := range s.words {
		s.words[i] &= op
	}
}

func (s *IntSet) Copy() *IntSet {
	n := len(s.words)
	words := make([]uint64, n)
	for i, word := range s.words {
		words[i] = word
	}
	return &IntSet{words}
}
