package queries

import "testing"

func TestBITQueryFunc(t *testing.T) {
	tests := []struct {
		q     BIT
		start int
		end   int
		sum   int
	}{
		{BIT{StaticArray: []int{1, 3, 8, 4, 6, 1, 3, 4}}, 3, 6, -3},
	}

	for i, tc := range tests {
		rq := tc.q
		sum, err := rq.sum(tc.start, tc.end)
		if err != nil {
			t.Fatalf("%s", err.Error())
		}
		if sum != tc.sum {
			t.Errorf("Test case %d expected sum value %d got %d", i, tc.sum, sum)
		}

	}

}
