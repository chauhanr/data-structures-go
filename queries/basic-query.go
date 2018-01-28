package queries

import (
	"errors"
	"math"
)

type RangeQuery interface {
	min(a, b int) (int, error)
	sum(a, b int) (int, error)
	max(a, b int) (int, error)
}

type BasicRangeQuery struct {
	StaticArray []int
	PrefixSum   []int
}

func (ba *BasicRangeQuery) Initialize() {
	sum := 0
	for _, val := range ba.StaticArray {
		sum += val
		ba.PrefixSum = append(ba.PrefixSum, sum)
	}
}

func (ba *BasicRangeQuery) min(a, b int) (int, error) {
	err := validate(a, b, ba.StaticArray)
	if err != nil {
		return 0, err
	}
	slice := ba.StaticArray[a : b+1]
	min := math.MaxInt16
	for _, val := range slice {
		if min > val {
			min = val
		}
	}

	return min, nil
}

func (ba *BasicRangeQuery) max(a, b int) (int, error) {
	err := validate(a, b, ba.StaticArray)
	if err != nil {
		return 0, err
	}
	slice := ba.StaticArray[a : b+1]
	max := math.MinInt16
	for _, val := range slice {
		if max < val {
			max = val
		}
	}
	return max, nil
}

func (ba *BasicRangeQuery) sum(a, b int) (int, error) {
	err := validate(a, b, ba.StaticArray)
	if err != nil {
		return 0, err
	}
	sum := 0
	sum = ba.PrefixSum[b] - ba.PrefixSum[a-1]
	return sum, nil
}

func validate(a int, b int, arr []int) error {
	if b < a {
		return errors.New("The end index cannot be greater than start")
	}
	if a < 0 || b >= len(arr) {
		return errors.New("invalid start and end index ")
	}
	return nil
}
