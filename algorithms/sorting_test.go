package algorithms

import (
	"testing"
)

var sortTests = []struct{
	arr []int64
	sortedArr []int64
}{
	{
		[]int64{9, 7, 6, 5, 2, 1},
		[]int64{1,2,5,6,7,9},
	},
	{
		[]int64{-9, 9, 3, -10, -12},
		[]int64{-12, -10, -9, 3, 9},
	},
	{
		[]int64{1, 2, 3, 9, 10},
		[]int64{1, 2, 3, 9, 10},
	},
}

func TestBubbleSort(t *testing.T){
	for _, tc := range sortTests {
		sortedResult := BubbleSort(tc.arr)
		for i, x :=range *sortedResult {
			y := tc.sortedArr[i]
			if x != y {
				t.Errorf("Values do not match expected %d but got %d", y, x)
			}
		}
	}
}

func TestMergeSort(t *testing.T){
	for _, tc := range sortTests {
		sortedResult := MergeSort(tc.arr, 0, len(tc.arr)-1)
		printArray(*sortedResult)
		for i, x :=range *sortedResult {
			y := tc.sortedArr[i]
			if x != y {
				t.Errorf("Values do not match expected %d but got %d", y, x)
			}
		}
	}
}

/*func printArray(arr []int64) {
	var strArr []string
	for _, element := range arr{
		strArr = append(strArr, strconv.FormatInt(element,10))
	}
	log.Printf("%s", strings.Join(strArr, ", "))

}*/

