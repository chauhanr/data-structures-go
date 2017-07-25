package queue

import (
	"fmt"
	"testing"
)

func TestQueueInit(t *testing.T) {
	q := Queue{}
	q.Initialize(10)
	size := q.QSize()
	if size != 0 {
		t.Errorf("Queue initialized with no value must have size 0 but was %d", size)
	}

	q.InitializeWithValue(8, 10)
	size = q.QSize()
	if size != 1 {
		t.Errorf("Queue initialized with no value must have size 1 but was %d", size)
	}
}

func TestZeroInitializationSize(t *testing.T) {
	q := Queue{}
	err := q.Initialize(0)
	if err == nil {
		t.Errorf("Queue should not allow for a queue to be initailized with a size <= 0")
	}
	err = q.InitializeWithValue(8, 0)
	if err == nil {
		t.Errorf("Queue should not allow for a queue to be initailized with a size <= 0")
	}
}

/**
This test case will test the following scenarios
1. it will enqueue 3 elements to a queue that has enough capacity.
2. then test the size of the queue is equal to the number of elements enqueued.
3. then iterate through the queue and check for the correct sequence.
4. finally test whether the queue is empty or not.
*/
func TestEnqueueFunc(t *testing.T) {
	q := Queue{}
	q.Initialize(10)
	buildQueue := [3]int{12, 10, 9}
	for _, value := range buildQueue {
		err := q.Enqueue(value)
		if err != nil {
			t.Errorf("Queue has enough room but error while enqueuing values")
		}
	}
	size := q.QSize()
	if size != 3 {
		t.Errorf("The queue size should be 3 but was %d", size)
	}

	for _, value := range buildQueue {
		qVal, err := q.Dequeue()
		if err != nil {
			t.Errorf("queue must have same number of elements as build queue")
		}
		if value != qVal {
			t.Errorf("Value from q should be %d but was %d", value, qVal)
		}
	}

	if !q.isEmpty() {
		t.Errorf("queue size should be 0 but was %d", q.fill)
	}

}

/**
This test cases tests the enqueue functionality when the queue is full.
The queue should return an error stating that the queue is full and you cannot add any more values
*/
func TestQueueFullFunctionality(t *testing.T) {
	q := Queue{}
	err := q.InitializeWithValue(3, 1)
	err = q.Enqueue(4)
	if err == nil {
		t.Errorf("Queue is was full but the Enqueue functions allowed another element to be added")
	} else {
		fmt.Printf("Error message : %s \n", err.Error())
	}
}

/**
This tests the dequeue functionality when the queue is empty.
The queue should not allow the client to dequeue from an empty queue.
*/
func TestEmptyQueueFunctionality(t *testing.T) {
	q := Queue{}
	q.Initialize(3)
	value, err := q.Dequeue()
	if err == nil {
		t.Errorf("Queue is suppose to have not value but returned %d", value)
	} else {
		fmt.Printf("%s\n", err.Error())
	}
}
