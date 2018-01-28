package amortized

import "testing"

func TestSubArraySum(t *testing.T) {
	tests := []struct {
		array []int
		sum   int
		start int
		end   int
	}{
		{
			[]int{1, 3, 2, 5, 1, 1, 2, 3},
			8,
			2,
			4,
		},
		{
			[]int{1, 3, 2, 5, 1, 1, 2, 3},
			6,
			0,
			2,
		},
		{
			[]int{1, 3, 2, 5, 1, 1, 2, 3},
			9,
			2,
			5,
		},
	}

	for i, tc := range tests {
		s, e := SubarraySum(tc.array, tc.sum)
		if s != tc.start && e != tc.end {
			t.Errorf("Test case %d expected [%d,%d] but got [%d,%d]", i, tc.start, tc.end, s, e)
		}
	}

}

func Test2Sum(t *testing.T) {
	tests := []struct {
		array []int
		value int
		start int
		end   int
	}{
		{
			[]int{1, 4, 5, 6, 7, 9, 9, 10},
			12,
			2,
			4,
		},
		{
			[]int{1, 4, 5, 6, 7, 9, 9, 10},
			20,
			-1,
			-1,
		},
	}

	for i, tc := range tests {
		s, e := Solve2Sum(tc.array, tc.value)
		if s != tc.start && e != tc.end {
			t.Errorf("Test case %d expected [%d,%d] but got [%d,%d]", i, tc.start, tc.end, s, e)
		}
	}
}
