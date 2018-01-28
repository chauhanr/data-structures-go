package dynamicp

import "testing"

func TestMaxGridSum(t *testing.T) {
	tests := []struct {
		grid   [][]int
		length int
		sum    int
	}{
		{
			[][]int{
				{3, 7, 9, 2, 7},
				{9, 8, 3, 5, 5},
				{1, 7, 9, 8, 5},
				{3, 8, 6, 4, 10},
				{6, 3, 9, 7, 8},
			},
			5,
			67,
		},
		{
			[][]int{
				{3, 7, 9, 2, 1},
				{9, 8, 3, 5, 5},
				{1, 7, 9, 8, 1},
				{3, 8, 6, 4, 1},
				{6, 3, 9, 7, 8},
			},
			5,
			66,
		},
	}

	for _, tc := range tests {
		max := MaxGridSum(tc.grid, tc.length)
		if max != tc.sum {
			t.Errorf("Expected sum %d got %d", tc.sum, max)
		}
	}
}
