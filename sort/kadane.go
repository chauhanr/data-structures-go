package sort

/**
	Kadane algorithm with give us the maximum sum of the sub array (without the order of the array being changed) in O(n) time.
    normal solution to the problem generally have O(n^3) or O(x^2) time complexity.
*/
func MaxSubArraySum(arr []int64) int64 {
	var best, sum int64
	best = 0
	sum = 0
	for _, e := range arr {
		sum = max(e, sum+e)
		best = max(best, sum)
	}
	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	} else {
		return b
	}

}
