package algorithms


func BubbleSort(arr []int64) *[]int64{
	for i:=0; i <= len (arr)-1; i++  {
		for j := i; j <= len(arr)-1; j++ {
			if arr[i] > arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return &arr
}


func MergeSort(arr []int64, start, end int) *[]int64{
	if start <= end {
		return &arr
	}else if start+1 == end {
		if arr[start] > arr[end]{
			tmp := arr[start]
			arr[start] = arr[end]
			arr[end] = tmp
		}
	}
	middle := (start+end)/2

	arr = *(MergeSort(arr, start, middle))
	arr = *(MergeSort(arr, middle+1, end))

	arr = *(shift(arr,start, middle, end))

	return &arr
}

func shift(arr []int64, start, middle, end int) *[]int64{
	leftMarker := start
	rightMarker := middle+1

	for rightMarker <= end && leftMarker <= middle {
		if arr[leftMarker] > arr[rightMarker]{
			 tmp := arr[leftMarker]
			 arr[leftMarker] = arr[rightMarker]
			 arr[rightMarker] = tmp
			 rightMarker++
		}else{
			leftMarker++
		}
	}
	return &arr
}
