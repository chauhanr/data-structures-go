package graphs

import (
	"math"
	"data-structures-go/queue"
	"container/heap"
)

type GraphMap struct {
	Alist map[int][]Element
	Nodes map[int]string
}

type Element struct {
	Dest   int
	Weight int
}

func (g *GraphMap) Initialize(){
	g.Alist = map[int][]Element{}
	g.Nodes = map[int]string{}
}

func (g *GraphMap) AddElement(src int, dest int, w int) {
	g.addUniqueNode(src)
	g.addUniqueNode(dest)

	e := Element{dest, w}
	if _, isPresent := g.Alist[src]; isPresent {
		elements := g.Alist[src]
		elements = append(elements, e)
		g.Alist[src] = elements
	} else {
		elements := []Element{e}
		g.Alist[src] = elements
	}
}

func (g *GraphMap) addUniqueNode(src int) {
	if _, isNodePresent := g.Nodes[src]; !isNodePresent{
		g.Nodes[src] = ""
	}
}

func (g *GraphMap) NodeCount() int{
	return len(g.Nodes)
}


func (g *GraphMap) Degree(node int) int{
	if _, isNodePresent := g.Alist[node]; isNodePresent{
		elements := g.Alist[node]
		return len(elements)
	}else {
		return 0
	}
}

func (g *GraphMap) DepthSearch(element int) []int{
	order := []int{}
	visited := map[int]string{}
	order = g.dfs(element,visited, order)
	return order
}

func (g *GraphMap) dfs(element int, visited map[int]string, order []int) []int{
	if _, isVisited := visited[element]; isVisited{
		return order
	}else {
		visited[element]=""
		order = append(order, element)
		neighbors := g.Alist[element]
		for _, lNode := range neighbors {
			order = g.dfs(lNode.Dest, visited, order)
		}
	}
	return order
}

func (g *GraphMap) BreathFirstSearch(element int) []int{
	order := map[int][]int{}
	visited := map[int]string{}
	g.bfs(element,visited, order, 1)
	result := []int{element}
	for i :=1; i<=len(order); i++{
		elements := order[i]
		result = append(result, elements...)
	}
	return result
}

func (g *GraphMap) bfs(element int, visited map[int]string, order map[int][]int, distance int) {
	if _, isVisited := visited[element]; isVisited{
		return
	}else {
		visited[element]=""
		orderElements := []int{}
		neighbors := g.Alist[element]
		for _, mNode := range neighbors{
			orderElements = append(orderElements, mNode.Dest)
		}
		if _, isExits := order[distance]; !isExits {
			order[distance] = orderElements
		}else {
			elements := order[distance]
			elements = mergeElements(elements, orderElements)
			order[distance] = elements
		}
		distance++
		for _, lNode := range neighbors {
			g.bfs(lNode.Dest, visited, order, distance)
		}
	}
	return
}

func mergeElements(a []int, b []int) []int{
	unique := map[int]string{}
	u := []int{}
	for _, v := range a {
		if _, exists := unique[v]; !exists{
			unique[v] = ""
			u = append(u, v)
		}
	}
	for _, v := range b{
		if _, exists := unique[v]; !exists{
			unique[v] = ""
			u = append(u, v)
		}
	}

	return u
}

func (g *GraphMap) Distance(nodeId int) []int{
	distance := make([]int, g.NodeCount()+1)
	// initialize to the infinity.
	for i:=0; i<=g.NodeCount(); i++{
		distance[i] = math.MaxInt32
	}
	nodeList := g.Alist[nodeId]
	distance[nodeId] = 0
	for _, node := range nodeList{
		distance = g.calculateDistance(nodeId, distance, node.Dest, node.Weight)
	}

	return distance
}

func (g *GraphMap) calculateDistance(src int, distance []int, dest int, w int) []int{
	// update the distance for this node.
	distance[dest] = min(distance[dest], distance[src]+ w)
	if nodeList, hasPath := g.Alist[dest]; hasPath{
		for _, node := range nodeList{
			distance = g.calculateDistance(dest, distance, node.Dest, node.Weight)
		}
	}
	return distance
}

func min(a, b int) int{
	if a > b {
		return b
	}else{
		return a
	}
}

func (g *GraphMap) DistanceDijkstra(nodeId int) []int{
	distance := make([]int, g.NodeCount()+1)
	processed := make([]bool, g.NodeCount()+1)

	for i:=0; i<len(distance); i++{
		distance[i] = math.MaxInt32
	}
	distance[nodeId] = 0
	pq := queue.PriorityQueue{}
	heap.Init(&pq)
	item := queue.Item{Value: nodeId, Priority: 0}
	heap.Push(&pq,&item)

	for pq.Len() > 0{
		top := heap.Pop(&pq).(*queue.Item)
		a := top.Value
		if processed[a] {
			continue
		}
		processed[a] = true
		for _, node := range g.Alist[a]{
			d := node.Dest
			w := node.Weight
			if distance[a]+w < distance[d]{
				distance[d] = distance[a]+w
				heapItem := queue.Item{Value: d, Priority: -distance[d]}
				heap.Push(&pq, &heapItem)
			}
		}
	}
	return distance
}