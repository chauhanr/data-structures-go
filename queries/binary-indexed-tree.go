package queries

type BIT struct {
	StaticArray []int
	Tree        []int
}

func (ba *BIT) Initialize() {

}

// give the sum of value from 1, k
func (ba *BIT) sumk(k int) int {
	sum := 0
	for k >= 1 {
		sum += ba.StaticArray[k]
		k -= k & (-k)
	}
	return sum
}

func (ba *BIT) sum(a, b int) (int, error) {
	err := validate(a, b, ba.StaticArray)
	if err != nil {
		return 0, err
	}

	sum := ba.sumk(b) - ba.sumk(a)
	return sum, nil
}
