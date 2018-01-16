package algorithms

import (
	"strconv"
	"log"
	"strings"
)

/**
	searches for element in an array using binary search algorithm.
	1. first sorts the arr and then
	2. searches for the element after the sort.
	3. return a bool for whether the element was found or not and then index of the element in the array.
	if the element is not found then the index is -1
*/
func BinarySearch(arr []int64, e int64) (bool, int){
	var middle, start, end int
	start = 0
	end = len(arr)-1

	for start <= end {
		middle = (start+end)/2
		if arr[middle] == e {
			return true, middle
		}else if arr[middle] > e {
			end = middle-1
		}else {
			start = middle+1
		}
	}
	return false, -1
}

func printArray(arr []int64) {
	var strArr []string
	for _, element := range arr{
		strArr = append(strArr, strconv.FormatInt(element,10))
	}
	log.Printf("%s", strings.Join(strArr, ", "))

}



