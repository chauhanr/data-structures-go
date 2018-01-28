package csearch

import (
	"testing"
)

func TestLeftDiagonalBuild(t *testing.T) {
	left := [][]int{
		{0, 1, 2, 3},
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
	}
	n := 4
	createdLeft := buildLeftDiagonal(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if createdLeft[i][j] != left[i][j] {
				t.Errorf("expeceted %d got %d at [%d,%d]", left[i][j], createdLeft[i][j], i, j)
			}
		}
	}
}

func TestRightDiagonalBuild(t *testing.T) {
	right := [][]int{
		{3, 4, 5, 6},
		{2, 3, 4, 5},
		{1, 2, 3, 4},
		{0, 1, 2, 3},
	}
	n := 4
	createdR := buildRightDiagonal(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if createdR[i][j] != right[i][j] {
				t.Errorf("expeceted %d got %d at [%d,%d]", right[i][j], createdR[i][j], i, j)
			}
		}
	}
}

func TestSolve8QueenProblem(t *testing.T) {
	solution := Solve8QueenProblem(5)

	t.Logf("Number of solutions for q queen for %d board is %d", 5, solution)
}
