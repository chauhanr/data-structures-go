package lists

import "errors"

//LinkedList is a structure of a node in the linked list.
type LinkedList struct {
	Head *Node
}

// Node represents the link on the linkedlist
type Node struct {
	Value    int64
	NextNode *Node
}

// Initialize method the linkedlist with head as nil
func (linkedList *LinkedList) Initialize() {
	linkedList.Head = nil
}

// InitializeWithValue method with the value will make a new head element.
func (linkedList *LinkedList) InitializeWithValue(value int64) {
	node := Node{}
	node.Value = value
	node.NextNode = nil
	linkedList.Head = &node
}

// size method will find the lenght of the list in questions.
func (linkedList *LinkedList) size() int64 {
	var size int64
	if linkedList.Head == nil {
		return size
	}
	ptr := linkedList.Head
	for ptr != nil {
		size++
		ptr = ptr.NextNode
	}
	return size
}

func (linkedList *LinkedList) search(item int64) *Node {
	if linkedList.Head == nil {
		return nil
	}
	ptr := linkedList.Head
	for ptr != nil {
		if ptr.Value == item {
			return ptr
		}
		ptr = ptr.NextNode
	}
	return nil
}

func (linkedList *LinkedList) delete(item int64) bool {
	searchNode := linkedList.search(item)
	if searchNode != nil {
		prev := linkedList.Head
		ptr := linkedList.Head
		if ptr == searchNode {
			linkedList.Head = ptr.NextNode
			return true
		}
		for ptr != searchNode {
			prev = ptr
			ptr = ptr.NextNode
			if ptr == searchNode {
				prev.NextNode = ptr.NextNode
				return true
			}
		}
	}
	return false
}

func (linkedList *LinkedList) append(value int64) {
	node := Node{}
	node.Value = value
	node.NextNode = nil

	if linkedList.Head == nil {
		linkedList.Head = &node
		return
	}
	prev := linkedList.Head
	ptr := linkedList.Head
	for ptr != nil {
		prev = ptr
		ptr = ptr.NextNode
	}
	prev.NextNode = &node
}

func (linkedList *LinkedList) average() float64 {
	var sum float64
	var size float64
	if linkedList.Head == nil {
		return 0.0
	}
	ptr := linkedList.Head
	for ptr != nil {
		size++
		sum = sum + float64(ptr.Value)
		ptr = ptr.NextNode
	}

	average := sum / size
	return average
}

// clear method will remove reference of the head and make it nil so the link list size is reduced to zero.
func (linkedList *LinkedList) clear() {
	if linkedList.Head == nil {
		// linked list is already clear and has size zero
		return
	}
	linkedList.Head = nil
}

/**
function will reverse the linkedlist in O(n) time
*/
func (linkedList *LinkedList) Reverse() error {
	if linkedList.Head == nil {
		return errors.New("Cannot reverse an empty list")
	}
	head := linkedList.Head
	linkedList.reverseList(head, head.NextNode)
	// final step make the next node for head a nil
	head.NextNode = nil
	return nil
}

func (linkedList *LinkedList) reverseList(parent *Node, current *Node) {
	if current == nil {
		return
	}
	linkedList.reverseList(current, current.NextNode)
	if current.NextNode == nil {
		linkedList.Head = current
	}
	current.NextNode = parent
}

func (linkedList *LinkedList) AsArray() []int64 {
	arr := []int64{}
	curr := linkedList.Head
	for curr != nil {
		arr = append(arr, curr.Value)
		curr = curr.NextNode
	}
	return arr
}
