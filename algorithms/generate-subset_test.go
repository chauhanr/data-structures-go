package algorithms

import (
	"testing"
	"log"
)

func TestGenerateSubset(t *testing.T){
	var subset = []int{}
	GenerateSubset(0, subset, 3)
	t.Logf("Completed the GenerateSubset Method...")
}


func TestGeneratePermutation(t *testing.T) {
	var perm = []int{}
	choosen := make([]bool, 3)
	log.Printf("%v", choosen)
	GeneratePermutation(perm, choosen , 3)
	t.Logf("Completed the GeneratePermutation Method...")
}