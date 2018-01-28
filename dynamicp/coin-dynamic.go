package dynamicp

import (
	"log"
	"math"
)

/**
function returns 0 in case of total = 0
returns -1 in case the total is < 0
return value in case total > 0
*/

func FindCoinCount(cs []int, total int) int {
	ready := make([]bool, total+1)
	value := make([]int, total+1)
	//countSol := make([]int, total+1)
	noOfCoins := solve(cs, total, ready, value)
	//solveI(cs,total)
	count := countSolutionI(make([]int, total+1), total, cs)
	//count := countSolution( total,cs)
	log.Printf("Solution Count %d ", count)
	return noOfCoins
}

func solveI(cs []int, total int) int {
	value := make([]int, total+1)
	value[0] = 0
	for x := 1; x < total; x++ {
		value[x] = math.MaxInt16
		for _, c := range cs {
			if x-c >= 0 {
				value[x] = min(value[x], value[x-c]+1)
			}
		}
	}
	//log.Printf("%v",value)
	return value[total-1]
}

func solve(cs []int, total int, ready []bool, value []int) int {
	best := math.MaxInt16
	if total == 0 {
		return 0
	}
	if total < 0 {
		return math.MaxInt16
	}
	if ready[total] {
		return value[total]
	}
	for _, coin := range cs {
		sol := solve(cs, total-coin, ready, value)
		best = min(best, sol+1)
	}
	ready[total] = true
	value[total] = best
	return best

}

func min(best int, solution int) int {
	if solution >= 0 {
		if best > solution {
			return solution
		} else {
			return best
		}
	}
	return solution
}

func countSolution(total int, cs []int) int {
	count := 0
	if total == 0 {
		return 1
	}
	if total <= 0 {
		return math.MaxInt16
	}
	for _, c := range cs {
		sol := countSolution(total-c, cs)
		if sol != math.MaxInt16 {
			count += sol
		}
	}
	return count
}

func countSolutionI(count []int, total int, cs []int) int {
	count[0] = 1
	for x := 1; x <= total; x++ {
		for _, c := range cs {
			if x-c >= 0 {
				count[x] += count[x-c]
			}
		}
	}
	log.Printf("%v", count)
	return count[total-1]
}
