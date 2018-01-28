package dynamicp

func MaxGridSum(grid [][]int, l int) int {
	sum := make([][]int, l)
	for i := range sum {
		sum[i] = make([]int, l)
	}

	copy(sum, grid, l)
	for x := 1; x <= l-1; x++ {
		for y := 1; y <= l-1; y++ {
			sum[x][y] = maxValue(sum[x-1][y], sum[x][y-1]) + grid[x][y]
		}
	}
	sum[l-1][l-1] += sum[0][0]
	return sum[l-1][l-1]
}

func copy(sum [][]int, grid [][]int, l int) {
	for x := 0; x < l; x++ {
		for y := 0; y < l; y++ {
			sum[x][y] = grid[x][y]
		}
	}
}
