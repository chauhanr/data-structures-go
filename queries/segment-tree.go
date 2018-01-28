package queries

import (
	"log"
	"math"
)

type SegmentTree struct {
	OrigArray []int
	SumTree   []int
	MinTree   []int
}

func (s *SegmentTree) Initialize() {
	lo := len(s.OrigArray)
	l := findTreeSize(len(s.OrigArray))
	s.SumTree = make([]int, l)
	s.MinTree = make([]int, l)

	j := l - 1
	for i := lo - 1; i > 0; {
		a := s.OrigArray[i]
		b := s.OrigArray[i-1]
		s.SumTree[j] = a + b
		s.MinTree[j] = minVal(a, b)
		i = i - 2
		j--
	}

	for j >= 0 {
		if j == 0 {
			s.SumTree[j] = s.SumTree[1] + s.SumTree[2]
			s.MinTree[j] = minVal(s.MinTree[1], s.MinTree[2])
		} else {
			s.SumTree[j] = s.SumTree[2*(j+1)] + s.SumTree[2*(j+1)-1]
			s.MinTree[j] = minVal(s.MinTree[2*(j+1)], s.MinTree[2*(j+1)-1])
		}
		j--
	}
	s.SumTree = append(s.SumTree, s.OrigArray...)
	s.MinTree = append(s.MinTree, s.OrigArray...)
	log.Printf("Min tree : %v", s.MinTree)
}

func (s *SegmentTree) sum(a int, b int) int {
	sum := 0
	a += len(s.OrigArray)
	b += len(s.OrigArray)

	for b >= a {
		if a%2 == 1 {
			sum += s.SumTree[a-1]
			a++
		}
		if b%2 == 0 {
			sum += s.SumTree[b-1]
			b--
		}
		a = a / 2
		b = b / 2
	}
	return sum
}

func (s *SegmentTree) min(a int, b int) int {
	min := math.MaxInt32
	a += len(s.OrigArray)
	b += len(s.OrigArray)

	for b >= a {
		if a%2 == 1 {
			min = minVal(min, s.MinTree[a-1])
			a++
		}
		if b%2 == 0 {
			min = minVal(min, s.MinTree[b-1])
			b--
		}
		a = a / 2
		b = b / 2
	}
	return min
}

func findTreeSize(l int) int {
	switch {
	case (l > 0 && l <= 8):
		return 7
	case l > 8 && l <= 16:
		return 15
	case l > 16 && l <= 32:
		return 31
	default:
		return 0
	}
}

func minVal(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
