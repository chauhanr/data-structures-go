package queue

import (
	"testing"
	"container/heap"
	"fmt"
)

func TestPriorityQueue(t *testing.T) {
	items := map[int]int{ 1: -3, 2: -2, 3: -4}
	pq := make(PriorityQueue, len(items))
	i := 0

	for value, priority := range items {
		//fmt.Printf("value : %d, priority: %d, index: %d \n", value, priority, i)
		pq[i] = &Item{value, priority,  i}
		i++
	}

	heap.Init(&pq)
	item := &Item{ Value : 4, Priority: -1}
	heap.Push(&pq, item)
	pq.update(item, item.Value, -5)

	for pq.Len() >0{
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("item : %d | priority : %d \n", item.Value, item.Priority)
	}

}
