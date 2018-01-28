package queries

import (
	"log"
	"testing"
)

func TestSegmentTreeInit(t *testing.T) {
	testsPrep := []struct {
		orig []int
		tree []int
	}{
		{[]int{5, 8, 6, 3, 2, 7, 2, 6}, []int{39, 22, 17, 13, 9, 9, 8, 5, 8, 6, 3, 2, 7, 2, 6}},
	}

	testSum := []struct {
		a   int
		b   int
		sum int
	}{
		{0, 7, 39},
		{0, 6, 33},
		{1, 6, 28},
		{0, 3, 22},
	}

	testMin := []struct {
		a   int
		b   int
		min int
	}{
		{0, 7, 2},
		{0, 3, 3},
	}

	for _, tc := range testsPrep {
		st := SegmentTree{OrigArray: tc.orig}
		st.Initialize()
		equal := areArrayEqual(st.SumTree, tc.tree)
		log.Printf("Tree %v: ", st.SumTree)
		if !equal {
			t.Errorf("Calculated Segment Tree %v is not equal to expected Tree %v", st.SumTree, tc.tree)
		}
		for _, ts := range testSum {
			sum := st.sum(ts.a, ts.b)
			if sum != ts.sum {
				t.Errorf("For sum(%d, %d) expected %d but got %d", ts.a, ts.b, ts.sum, sum)
			}
		}
		for _, tm := range testMin {
			min := st.min(tm.a, tm.b)
			if min != tm.min {
				t.Errorf("For min(%d, %d) expected %d but got %d", tm.a, tm.b, tm.min, min)
			}
		}
	}
}

func areArrayEqual(orig []int, tree []int) bool {
	for i := 0; i < len(orig); i++ {
		if orig[i] != tree[i] {
			return false
		}
	}
	return true
}
