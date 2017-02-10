package main

import (
	"bytes"
	"fmt"
)

/*
* 实现一个类似set的Bit数组,方便查找
 */

type Inset struct {
	words []uint64
}

func (s *Inset) Add(x int) {
	word, index := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << index
}

func (s *Inset) Has(x int) bool {
	word, index := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<index) != 0
}

func (s *Inset) UnionSet(a *Inset) {
	for i, word := range a.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *Inset) String() string {
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
				value := i*64 + j
				fmt.Fprintf(&buf, "%d", value)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var s, y Inset
	s.Add(5)
	s.Add(10)
	s.Add(120)
	fmt.Println(s.Has(10))
	fmt.Println(s.String())
	y.Add(2)
	y.Add(8)
	y.Add(11)
	s.UnionSet(&y)
	fmt.Println(s.String())
}
