package algorithms

import "testing"

func TestKadaneMaxSubset(t *testing.T){
	tests := []struct{
		array []int64
		maxSum int64
	}{
		{
			[]int64{-1, 2, 4, -3, 5, 2, -5, 2},
			10,
		},
		{
			[]int64{-1, -2, 4, -3, 5, 2, -5, 2, 7},
			12,
		},
	}

	for _, tc := range tests{
		best := MaxSubArraySum(tc.array)
		if best != tc.maxSum{
			t.Errorf("Expected the max sum to be %d but got %d",tc.maxSum,best)
		}
	}
}
