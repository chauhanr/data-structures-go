package dynamicp

import "testing"

func TestKnapsackProblem(t *testing.T) {
	tests := []struct {
		input  []int
		output []bool
	}{
		{
			[]int{1, 3, 3, 5},
			[]bool{true, true, false, true, true, true, true, true, true, true, false, true, true},
		},
	}

	for _, tc := range tests {
		result := SolveKnapsack(len(tc.output), tc.input)
		t.Logf("%v", result)
		/*for index, b := range result {
			if b != tc.output[index] {
				t.Errorf("For sum %d expected %v but got %v", index, tc.output[index], b)
			}
		}*/
	}
}
