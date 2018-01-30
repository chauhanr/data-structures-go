package graphs

import (
	"testing"
	"strconv"
	"strings"
	"math"
	"log"
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

func TestBreathFirstSearch(t *testing.T){
	testBFS := []struct{
		element int
		order []int
	}{
		{1, []int{1, 2,4,3,5,6}},
	}

	graph := prepareBFSGraph()

	for _, tc := range testBFS {
		result := graph.BreathFirstSearch(tc.element)
		orderStr := stringify(tc.order)
		resultStr := stringify(result)
		if orderStr != resultStr{
			t.Errorf("Traversal should be %s but got %s", orderStr, resultStr)
		}
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

func TestGraphMap_Distance(t *testing.T) {
	testBFAlgo := []struct{
		nodeId int
		distance []int
	}{
		{1, []int{math.MaxInt32,0,2,3,1,3}},
		{2, []int{math.MaxInt32,math.MaxInt32,0,math.MaxInt32, 3, 5}},
	}

	graph := prepareBellFordAlgoGraph()

	for _, tc := range testBFAlgo{
		distArr := graph.Distance(tc.nodeId)
		if len(tc.distance) != len(distArr) {
			t.Fatalf("Expected the distance to be %v got %v the are of unequal length", tc.distance, distArr)
		}
		log.Printf("distance %v", distArr)
		for i, val := range distArr{
			if val != tc.distance[i]{
				t.Errorf("Expected distance at node %d is %d got %d", i, tc.distance[i], val)
			}
		}
	}
}

func TestGraphMap_DistanceDijkstra(t *testing.T) {
	testDAlgo := []struct{
		nodeId int
		distance []int
	}{
		{1, []int{math.MaxInt32,0,5,7,3,1}},
	}

	graph := prepareDijstrasAlgoGraph()

	for _, tc := range testDAlgo{
		distArr := graph.DistanceDijkstra(tc.nodeId)
		if len(tc.distance) != len(distArr) {
			t.Fatalf("Expected the distance to be %v got %v the are of unequal length", tc.distance, distArr)
		}
		log.Printf("distance %v", distArr)
		for i, val := range distArr{
			if val != tc.distance[i]{
				t.Errorf("Expected distance at node %d is %d got %d", i, tc.distance[i], val)
			}
		}
	}

}

func prepareSimpleGraph() GraphMap{
	graph := GraphMap{}
	graph.Initialize()
	graph.AddElement(1, 2, 5)
	graph.AddElement(2,3, 7)
	graph.AddElement(2,4,6)
	graph.AddElement(3,4,5)
	graph.AddElement(4,1,2)
	return graph
}

func prepareDFSGraph() GraphMap{
	graph := GraphMap{}
	graph.Initialize()
	graph.AddElement(1,2,1)
	graph.AddElement(2,3,1)
	graph.AddElement(3,5,1)
	graph.AddElement(2,5,1)
	graph.AddElement(1,4,1)
	return graph
}

func prepareBFSGraph() GraphMap{
	graph := GraphMap{}
	graph.Initialize()
	graph.AddElement(1,2,1)
	graph.AddElement(2,3,1)
	graph.AddElement(3,6,1)
	graph.AddElement(2,5,1)
	graph.AddElement(1,4,1)
	graph.AddElement(5,6,1)
	return graph
}

func prepareBellFordAlgoGraph() GraphMap{
	graph := GraphMap{}
	graph.Initialize()
	graph.AddElement(1,2,2)
	graph.AddElement(1,3,3)
	graph.AddElement(1,4, 7)
	graph.AddElement(2, 5, 6)
	graph.AddElement(2,4, 3)
	graph.AddElement(3,4, -2)
	graph.AddElement(4,5 , 2)
	return graph
}

func prepareDijstrasAlgoGraph() GraphMap{
	graph := GraphMap{}
	graph.Initialize()
	graph.AddElement(1,5,1)
	graph.AddElement(1,4,9)
	graph.AddElement(1,2, 5)
	graph.AddElement(2, 3, 2)
	graph.AddElement(5,4, 2)
	graph.AddElement(4,3 , 6)
	return graph
}