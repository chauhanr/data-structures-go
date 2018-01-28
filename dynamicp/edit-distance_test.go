package dynamicp

import "testing"

func TestEditDistance(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		distance int
	}{
		{"love", "movie", 2},
	}

	for i, tc := range tests {
		d := EditDistance(tc.a, tc.b)
		if d != tc.distance {
			t.Errorf("Expected distance for test case %d is %d but got %d", i, tc.distance, d)
		}
	}

}
