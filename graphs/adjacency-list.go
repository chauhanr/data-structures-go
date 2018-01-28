package graphs


type GraphMap struct {
	Alist map[int][]Element
}

type Element struct {
	Dest   int
	Weight int
}

func (g *GraphMap) Inititialize() {
	g.Alist = make(map[int][]Element)
}

func (g *GraphMap) AddElement(src int, dest int, w int) {
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

func (g *GraphMap) NodeCount() int{
	nodes := len(g.Alist)
	return nodes
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