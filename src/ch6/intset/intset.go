// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

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
func (s *IntSet) UnionWith(t *IntSet) {
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

//Ex6.1, return the number of elements
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count += 1
			}
		}
	}
	return count
}

//Ex6.1, remove x from the set
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}

	word, bit := x/64, uint(x%64)
	s.words[word] ^= 1 << bit //XOR
	return
}

//Ex6.1, remove all elements from the set
func (s *IntSet) Clear() {
	for i := 0; i < len(s.words); i++ {
		s.words[i] = 0
	}
	return
}

//Ex6.1, return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var t IntSet
	wordSize := len(s.words)

	twords := make([]uint64, wordSize)
	for i := 0; i < wordSize; i++ {
		twords[i] = s.words[i]
	}
	t.words = twords

	return &t
}

//Ex6.2, a variadic AddAll funtion that allows a list of values to be added
func (s *IntSet) AddAll(values ...int) {
	for _, value := range values {
		s.Add(value)
	}
}

//Ex6.3, insersection between two intsets
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i <= len(s.words) {
			s.words[i] &= tword
		}
	}
}

//Ex6.3, elements in S but not in T
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i <= len(s.words) {
			s.words[i] &= ^tword // Difference = x & (^y)
		}
	}
}

//Ex6.3, elements ONLY in S or ONLY in T
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i <= len(s.words) {
			s.words[i] ^= tword // Difference = x XOR y
		}
	}
}
