package algorithms

import "testing"

func TestFindCoinCount(t *testing.T) {
	tests := []struct{
		id int
		cs []int
		total int
		result int
	}{
		{
			1,
			[]int{1,3,4},
			10,
			3,
		},
		{
			2,
			[]int{1,3,4},
			3,
			1,
		},{
			3,
			[]int{1,3,4},
			5,
			2,
		},{
			4,
			[]int{1,3,4},
			2,
			2,
		},
		{
			5,
			[]int{1,3,4},
			9,
			3,
		},{
			6,
			[]int{1,2,3,5,10,20,50,100},
			540,
			 7,
		},
	}

	for _, tc := range tests{

		solution := FindCoinCount(tc.cs,tc.total)
		if solution != tc.result {
			t.Errorf("Test case %d failed expected %d got %d",tc.id, tc.result, solution)
		}
	}

}
