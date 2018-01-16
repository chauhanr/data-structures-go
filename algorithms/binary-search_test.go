package algorithms

import "testing"

var searchTests = []struct{
	arr []int64
	element int64
	found bool
	index int
}{
	{
		[]int64{1, 2, 5, 6, 7, 9},
		1,
		true,
		0,
	},
	{
		[]int64{-12, -10, -9, 3, 9},
		9,
		 true,
		 4,
	},
	{
		[]int64{-12, -10, -9, 3, 9},
		2,
		false,
		-1,
	},
}


func TestBinarySearch(t *testing.T){

	for _, tc := range searchTests{
		found, index := BinarySearch(tc.arr,tc.element)
		if found != tc.found || index != tc.index{
			t.Errorf("Expected to find element %d at index %d", tc.element, index)
		}
	}


}
