package algorithms

import "testing"

func TestFindLongestSunsequence(t *testing.T) {
	tests := []struct{
		id int
		cs []int
		result int
	}{
		{
			1,
			[]int{6,2,5,1,7,4,8,3},
			4,
		},
		{
			2,
			[]int{1,3,4},
			3,
		},

	}

	for _, tc := range tests{
		length := make([]int, len(tc.cs))
		FindLongestSubsequence(tc.cs,length)
		t.Logf("%v",length)
	}
}
