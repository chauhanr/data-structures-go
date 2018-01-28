package dynamicp

import "log"

func SolveKnapsack(max int, w []int) []bool {
	returnArr := make([]bool, max)
	possible := make([][]bool, max)
	for p := range possible {
		possible[p] = make([]bool, len(w)+1)
	}
	possible[0][0] = true
	for k := 1; k <= len(w); k++ {
		for x := 0; x < max; x++ {
			if k < len(w) && x-w[k] >= 0 {
				possible[x][k] = possible[x][k] || possible[x-w[k]][k-1]
			}
			possible[x][k] = possible[x][k] || possible[x][k-1]
			returnArr[x] = returnArr[x] || possible[x][k]
		}
	}
	log.Printf("%v", possible)
	return returnArr

}
