package queue

import "errors"

type Queue struct {
	Head *QNode
	Tail *QNode
	Size int
	fill int
}

type QNode struct {
	Value    int
	NextNode *QNode
	PrevNode *QNode
}

func (queue *Queue) Initialize(size int) error {
	if size <= 0 {
		return errors.New("Cannot have a queue fo size zero or less.")
	}
	queue.Head = nil
	queue.Tail = nil
	queue.fill = 0
	queue.Size = size
	return nil
}

func (queue *Queue) InitializeWithValue(value int, size int) error {
	if size <= 0 {
		return errors.New("Cannot have a queue fo size zero or less.")
	}
	node := QNode{}
	node.Value = value
	node.NextNode = nil
	node.PrevNode = nil
	queue.Head = &node
	queue.Tail = queue.Head
	queue.fill++
	queue.Size = size
	return nil
}

/**
The enquque method will add an element but only if there is space in the queue.
The enqueue method will add to the head of the queue.
*/
func (queue *Queue) Enqueue(value int) error {
	if queue.isFull() {
		return errors.New("Queue is full cannot enter any more values. ")
	}
	node := QNode{}
	node.Value = value
	if queue.Head == nil {
		queue.Head = &node
		queue.Tail = &node
	} else {
		prevHead := queue.Head
		node.NextNode = prevHead
		queue.Head = &node
		prevHead.PrevNode = queue.Head
	}
	queue.fill++
	return nil
}

/**
Simple method that will test of the queue is empty or not.
*/
func (queue *Queue) isEmpty() bool {
	if queue.fill == 0 {
		return true
	}
	return false
}

func (queue *Queue) isFull() bool {
	if queue.fill == queue.Size {
		return true
	}
	return false
}

/**
Dequeueu method will take the tail element and then remove it.
it will send the second last element in the queue as the tail.
*/
func (queue *Queue) Dequeue() (int, error) {
	if queue.isEmpty() {
		return -1, errors.New("Cannot dequeue from empty queue")
	}
	tail := queue.Tail
	newTail := tail.PrevNode
	queue.Tail = newTail
	if newTail != nil {
		newTail.NextNode = nil
	}
	queue.fill--
	return tail.Value, nil
}

/**
returns the fill value of the queue
*/
func (queue *Queue) QSize() int {
	return queue.fill
}
