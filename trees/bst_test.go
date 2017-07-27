package trees

import (
	"fmt"
	"testing"
)

/**
  Test the Initialize function of the BST.
*/
func TestBSTInitializeFunction(t *testing.T) {
	tree := BinaryTree{}
	tree.Initialize()
	if tree.Size() != 0 {
		t.Errorf("Initialized Tree should have size 0 but was %d", tree.Size())
	}
}

var testInsertStruct = []struct {
	elements []int
	result   []int
}{
	{[]int{3, 7, 8, 1, 9, 2}, []int{1, 2, 3, 7, 8, 9}},
}

func TestBSTInsertFunc(t *testing.T) {
	tree := BinaryTree{}
	tree.Initialize()
	for _, testCases := range testInsertStruct {
		buildTree := testCases.elements
		for _, value := range buildTree {
			tree.Insert(value)
		}
		size := tree.Size()
		if size != len(buildTree) {
			t.Errorf("Expected size of the tree is 6 but value returned is %d", size)
		}
		expectedResult := testCases.result

		traversed := []int{}
		traversed = tree.inOrderTraveral(tree.Root, traversed)

		for index, value := range traversed {
			if expectedResult[index] != value {
				t.Errorf("Value at index %d was supposed to be %d but was %d", index, expectedResult[index], value)
			}
		}
	}
}

// search for an elemetn in the collection and another without
func TestBSTSearch(t *testing.T) {
	tree := BinaryTree{}
	tree.Initialize()

	// search in an empty tree should give no result.
	node := tree.SearchItem(9)
	if node != nil {
		t.Errorf("Tree should not have value 9 but was found")
	}

	buildTree := [6]int{7, 3, 1, 8, 9, 2}
	for _, value := range buildTree {
		tree.Insert(value)
	}

	node = tree.SearchItem(1)
	if node == nil {
		t.Errorf("Tree should have value 1 but was not found")
	}

	node = tree.SearchItem(7)
	if node == nil {
		t.Errorf("Tree should have value 7 but was not found")
	}

	node = tree.SearchItem(9)
	if node == nil {
		t.Errorf("Tree should have value 9 but was not found")
	}

	node = tree.SearchItem(11)
	if node != nil {
		t.Errorf("Tree should not have value 11 but a node %d was found", node.Item)
	}
}

var testDeleteStruct = []struct {
	elements   []int
	deleteItem int
	result     []int
}{
	{[]int{3, 7, 9, 8, 1, 2, 10}, 2, []int{1, 3, 7, 8, 9, 10}},
	{[]int{3, 7, 9, 8, 1, 2, 10}, 7, []int{1, 2, 3, 8, 9, 10}},
	{[]int{3, 4, 9, 8, 1, 2, 10}, 9, []int{1, 2, 3, 4, 8, 10}},
	{[]int{6, 3, 9, 8, 7, 1, 10}, 6, []int{1, 3, 7, 8, 9, 10}},
	{[]int{6, 3, 9, 8, 7, 1, 10}, 8, []int{1, 3, 6, 7, 9, 10}},
}

func TestBSTDeleteNodeFunc(t *testing.T) {
	for _, testCases := range testDeleteStruct {
		tree := BinaryTree{}
		tree.Initialize()
		buildTree := testCases.elements
		for _, value := range buildTree {
			tree.Insert(value)
		}

		tree.DeleteNodeByValue(testCases.deleteItem)
		traversed := []int{}
		traversed = tree.inOrderTraveral(tree.Root, traversed)
		expectedResult := testCases.result
		fmt.Printf("%v\n", traversed)
		for index, value := range traversed {
			if expectedResult[index] != value {
				t.Errorf("Value at index %d was supposed to be %d but was %d", index, expectedResult[index], value)
			}
		}
	}
}

var testMinMaxStruct = []struct {
	elements []int
	min      int
	max      int
}{
	{[]int{3, 7, 9, 8, 2, 10}, 2, 10},
	{[]int{3, 7, 9, 8, 5, 6, 4}, 3, 9},
}

func TestMinMaxFunction(t *testing.T) {
	tree := BinaryTree{}
	tree.Initialize()

	_, err := tree.FindMinimumElement()
	if err == nil {
		t.Errorf("Empty tree cannot have a min/max element")
	}

	_, err = tree.FindMaxElement()
	if err == nil {
		t.Errorf("Empty tree cannot have a min/max element")
	}

	for _, testCase := range testMinMaxStruct {
		tree = BinaryTree{}
		tree.Initialize()
		buildTree := testCase.elements
		minResult := testCase.min
		maxResult := testCase.max
		for _, value := range buildTree {
			tree.Insert(value)
		}
		min, _ := tree.FindMinimumElement()
		if min != minResult {
			t.Errorf("Minimum element should be %d by the function returned %d", minResult, min)
		}
		max, _ := tree.FindMaxElement()
		if max != maxResult {
			t.Errorf("Maximum element should be %d by the function returned %d", maxResult, max)
		}
	}
}

var testTreeCompareStruct = []struct {
	tree1          []int
	tree2          []int
	expectedResult bool
}{
	{[]int{3, 7, 9, 8, 2, 10}, []int{3, 7, 9, 8, 2, 10}, true},
	{[]int{4, 7, 9, 8, 2, 10}, []int{4, 7, 9, 8, 3, 10}, false},
	{[]int{3, 7, 9, 8, 2, 10}, []int{7, 3, 8, 9, 10, 2}, false},
	{[]int{4, 7, 9, 8, 2, 10}, []int{4, 7, 9, 8, 2, 11}, false},
	{[]int{3, 7, 9, 8, 2, 10}, []int{7}, false},
	{[]int{}, []int{}, true},
}

func TestTreeCompareFunction(t *testing.T) {

	for _, testCompareCase := range testTreeCompareStruct {
		case1 := testCompareCase.tree1
		case2 := testCompareCase.tree2
		tree1 := BinaryTree{}
		tree1.Initialize()
		for _, v1 := range case1 {
			tree1.Insert(v1)
		}
		tree2 := BinaryTree{}
		tree2.Initialize()
		for _, v2 := range case2 {
			tree2.Insert(v2)
		}
		result := CompareBSTs(&tree1, &tree2)
		eResult := testCompareCase.expectedResult

		if result != eResult {
			if result == false {
				t.Errorf("Trees %v and %v  are equal but result of comparison was false", case1, case2)
			} else {
				t.Errorf("Trees %v and %v are not equal but result of comparison was true", case1, case2)
			}
		}
	}
}
