package algorithms

import (
	"math"
	"log"
)

/**
	function returns 0 in case of total = 0
	returns -1 in case the total is < 0
	return value in case total > 0
*/

func  FindCoinCount(cs []int, total int) int{
	ready := make([]bool,total+1)
	value := make([]int,total+1)
	noOfCoins := solve(cs, total, ready, value)
	solCount := make([]int, total+1)
	count := countSolution(solCount,total, cs)
	log.Printf("%v", count)
	//solveI(cs,total)
	return noOfCoins
}

func solveI(cs []int, total int) int{
	value := make([]int, total+1)
	value[0] = 0
	for x := 1; x <total; x++{
		value[x] = math.MaxInt16
		for _, c := range cs{
			if x -c >= 0{
				value[x] = min(value[x],value[x-c]+1)
			}
		}
	}
	//log.Printf("%v",value)
	return value[total-1]
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

func countSolution(count []int, total int, cs []int) []int{
	count[0] = 1
	for x := 1; x <total; x++{
		for _, coin := range cs {
			if x - coin >= 0 {
				count[x] += count[x-coin]
			}
		}
	}
	return count
}