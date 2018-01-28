package queries

import "testing"

func TestBasicQueryFunc(t *testing.T) {
	tests := []struct {
		q     BasicRangeQuery
		start int
		end   int
		min   int
		max   int
		sum   int
	}{
		{
			BasicRangeQuery{StaticArray: []int{1, 3, 8, 4, 6, 1, 3, 4}},
			3,
			6,
			1,
			6,
			14,
		},
	}

	for i, tc := range tests {
		rq := tc.q
		rq.Initialize()
		min, _ := rq.min(tc.start, tc.end)
		max, _ := rq.max(tc.start, tc.end)
		sum, _ := rq.sum(tc.start, tc.end)

		if min != tc.min {
			t.Errorf("Test case %d expected min value %d got %d", i, tc.min, min)
		}
		if max != tc.max {
			t.Errorf("Test case %d expected max value %d got %d", i, tc.max, max)
		}
		if sum != tc.sum {
			t.Errorf("Test case %d expected sum value %d got %d", i, tc.sum, sum)
		}

	}

}
