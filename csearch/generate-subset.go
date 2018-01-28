package csearch

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func GenerateSubset(k int, subset []int, n int) {
	if k == n {
		printSubset(subset)
	} else {
		GenerateSubset(k+1, subset, n)
		subset = append(subset, k)
		GenerateSubset(k+1, subset, n)
		if subset != nil && len(subset) > 1 {
			if len(subset) <= 1 {
				subset = []int{}
			} else if len(subset) == 2 {
				subset = append([]int{}, subset[0])
			} else {
				subset = subset[0 : len(subset)-2]
			}
		}
	}
}

func GeneratePermutation(permutation []int, choosen []bool, n int) error {
	if len(choosen) != n || permutation == nil {
		return errors.New("wrong operating parameters")
	}
	if len(permutation) == n {
		// process permutation
		printSubset(permutation)
	} else {
		for i := 0; i < n; i++ {
			if choosen[i] {
				continue
			}
			choosen[i] = true
			permutation = append(permutation, i)
			GeneratePermutation(permutation, choosen, n)
			choosen[i] = false
			if len(permutation) <= 1 {
				permutation = []int{}
			} else if len(permutation) == 2 {
				permutation = append([]int{}, permutation[0])
			} else {
				permutation = permutation[0 : len(permutation)-2]
			}
		}
	}
	return nil
}

func printSubset(subset []int) {
	var strArr []string
	if subset != nil && len(subset) != 0 {
		for _, element := range subset {
			strArr = append(strArr, ""+strconv.Itoa(element))
		}
		log.Printf("%s", strings.Join(strArr, ", "))
	} else {
		log.Printf("null")
	}

}
