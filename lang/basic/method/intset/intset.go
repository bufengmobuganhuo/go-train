package intset

import (
	"bytes"
	"fmt"
)

const BIT_SIZE = 64

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	// 对于一个数x，i=x/64, j = x%64
	// 如果words[i]的第j位=1，则说明x存在
	wordIdx, bitIdx := x/BIT_SIZE, x%BIT_SIZE
	for wordIdx >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// 将对应位置为1
	s.words[wordIdx] |= 1 << bitIdx
}

func (s *IntSet) Has(x int) bool {
	wordIdx, bitIdx := x/BIT_SIZE, x%BIT_SIZE
	return wordIdx < len(s.words) && s.words[wordIdx]&(1<<bitIdx) != 0
}

// 取并集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tWord := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, tWord)
		} else {
			s.words[i] |= tWord
		}
	}
}

// 返回集合中元素个数
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) Remove(x int) {
	wordIdx, bitIdx := x/BIT_SIZE, x%BIT_SIZE
	if wordIdx >= len(s.words) {
		return
	}
	s.words[wordIdx] = s.words[wordIdx] ^ (1 << bitIdx)
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		// word=0，说明每一位都是0，没有元素
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			// 说明对应位有值
			if word&(1<<j) != 0 {
				// 之前已经加过元素
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				// 恢复原来的数
				x := i*64 + j
				fmt.Fprintf(&buf, "%d", x)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
