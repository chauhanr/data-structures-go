package graphs

import (
	"testing"
	"strconv"
	"strings"
)

func TestGraphMapNodeCount(t *testing.T){
	graph := prepareSimpleGraph()
	if graph.NodeCount() != 4 {
		t.Errorf("Node count expected to be %d got %d", 4, graph.NodeCount())
	}
	degree2 := graph.Degree(2)
	if degree2 != 2 {
		t.Errorf("Degree for node 2 is expected to be 2 but found %d", degree2)
	}
	degree4 := graph.Degree(4)
	if degree4 != 1 {
		t.Errorf("Degree for node 4 is expected to be 1 but found %d", degree4)
	}
}

func TestDepthFirstSearch(t *testing.T){
	testDFS := []struct{
		element int
		order []int
	}{
		{1, []int{1,2,3,5,4}},
		{3, []int{3,5}},
		{4, []int{4}},
	}

	graph := prepareDFSGraph()

	for _, tc := range testDFS {
		result := graph.DepthSearch(tc.element)
		orderStr := stringify(tc.order)
		resultStr := stringify(result)
		if orderStr != resultStr{
			t.Errorf("Traversal should be %s but got %s", orderStr, resultStr)
		}
	}
}

func stringify(array []int) string{
	str := []string{}
	for _, val := range array{
		str = append(str, strconv.Itoa(val))
	}
	return strings.Join(str,",")
}

func prepareSimpleGraph() GraphMap{
	graph := GraphMap{}
	graph.Inititialize()
	graph.AddElement(1, 2, 5)
	graph.AddElement(2,3, 7)
	graph.AddElement(2,4,6)
	graph.AddElement(3,4,5)
	graph.AddElement(4,1,2)
	return graph
}

func prepareDFSGraph() GraphMap{
	graph := GraphMap{}
	graph.Inititialize()
	graph.AddElement(1,2,1)
	graph.AddElement(2,3,1)
	graph.AddElement(3,5,1)
	graph.AddElement(2,5,1)
	graph.AddElement(1,4,1)
	return graph
}
