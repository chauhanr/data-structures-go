package algorithms

import (
	"math"
)

/**
	function returns 0 in case of total = 0
	returns -1 in case the total is < 0
	return value in case total > 0
*/

func  FindCoinCount(cs []int, total int) int{
	ready := make([]bool,total+1)
	value := make([]int,total+1)
	count := solve(cs, total, ready,value)
	return count
}

func  solve(cs []int, total int, ready []bool, value []int) int{
	best := math.MaxInt16
	if total == 0{
		return 0
	}
	if total < 0{
		return math.MaxInt16
	}
	if ready[total] {
		return value[total]
	}
	for _, coin := range cs{
		best = min(best,solve(cs,total-coin, ready, value)+1)
	}
	ready[total] = true
	value[total] = best
	return best

}

func min(best int, solution int) int{
	if solution >= 0 {
		if best > solution{
			return solution
		}else{
			return best
		}
	}
	return solution
}



