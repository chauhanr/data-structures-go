package linkedList

import (
	"fmt"
	"testing"
)

func TestLinkedListInit(t *testing.T) {
	ll := LinkedList{}
	ll.Initialize()
	size := ll.size()
	if size != 0 {
		t.Errorf("LinkedList Initilaized with no value should be of size %d but was %d \n", 0, size)
	}

	ll.InitializeWithValue(8)
	size = ll.size()
	if size != 1 {
		t.Errorf("LinkedList Initialized with value should be of size %d but was %d \n", 1, size)
	}
}

func TestAppendFunc(t *testing.T) {
	ll := LinkedList{}
	ll.Initialize()
	buildList := [3]int64{13, 19, 21}
	for _, val := range buildList {
		ll.append(val)
	}
	size := ll.size()
	eSize := int64(len(buildList))
	if size != eSize {
		t.Errorf("Expected Length was %d but actual length was %d ", eSize, size)
	}
}

func TestClearFunc(t *testing.T) {
	ll := LinkedList{}
	ll.Initialize()
	buildList := [3]int64{13, 19, 21}
	for _, val := range buildList {
		ll.append(val)
	}
	ll.clear()
	size := ll.size()
	if size != 0 {
		t.Errorf("Clear should reduce the size of the list to 0 but the size is %d", size)
	}
}

func TestClearEmptyFunc(t *testing.T) {
	ll := LinkedList{}
	ll.Initialize()
	ll.clear()
	size := ll.size()
	if size != 0 {
		t.Errorf("Clear should reduce the size of the list to 0 but the size is %d", size)
	}
}

var testAverageStruct = []struct {
	elements []int64
	average  float64
}{
	{[]int64{21, 31, 13, 15, 14}, 18.8},
	{[]int64{3, 9, 4, 21, 4}, 8.2},
}

/**
This test case will add a few links into the list and then calculate the average
*/
func TestLinkedListAverageFunc(t *testing.T) {
	ll := LinkedList{}
	ll.Initialize()
	for _, avgTest := range testAverageStruct {
		el := avgTest.elements
		for _, val := range el {
			ll.append(val)
		}
		avg := ll.average()
		fmt.Printf("average for buildList %v , is %f\n", el, avg)
		if avg != avgTest.average {
			t.Errorf("Expected average is %f but got %f", avgTest.average, avg)
		}
		ll.clear()
	}
}

/**
This method tests the size of the linkedlist that is zero and has been initialized as zero.
*/
func TestLinkedListAverageZeroSize(t *testing.T) {
	ll := LinkedList{}
	ll.Initialize()
	avg := ll.average()
	if avg != 0.0 {
		t.Errorf("Expected average for a zero linked list is %f, but found %f", 0.0, avg)
	}
}
