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
}

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
