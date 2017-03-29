package linkedList

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
