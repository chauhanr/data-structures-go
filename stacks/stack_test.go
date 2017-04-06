package stacks

import "testing"

var stackEle = []int{2, 45, 99, 108, 32, 1, 7, 8}

/**
  pushing an element into an un instantiated stack must give an error
*/
func TestStackInstantiationSimple(t *testing.T) {
	stack := ArrayStack{}
	element := 1
	err := stack.Push(element)
	if err == nil {
		t.Errorf("Trying to push element in uninstantiated array should give an error")
	} else {
		t.Logf("Error in Pushing element in uninstantiated stack test case pass")
	}
}

/**
  Test case to check of the initialization fails when done with a negative value.
*/
func TestInstantiationWithNegSize(t *testing.T) {
	stack := ArrayStack{}
	err := stack.NewArrayStack(-2)
	if err == nil {
		t.Errorf("There should be an error when instantiating an array with negative size but was not | test FAIL")
	} else {
		t.Logf("Error while instantiating stack with size < 0 | test PASS")
	}
}

/**
  Test case to check of the initialization fails when done with a positive value.
*/
func TestStackInstantiationPosValues(t *testing.T) {
	stack := ArrayStack{}
	element := 1
	err := stack.NewArrayStack(32)
	if err != nil {
		t.Errorf("Instantiation must not fail when the length of the array is > 0")
	}
	// now the array is fine and has been instantiated. should not give an error
	err = stack.Push(element)
	if err != nil {
		t.Errorf("Should not get an error pushing element to stack as it is instantiated")
	} else {
		if stack.Fill() == 1 {
			t.Logf("The Fill of the stack is : %d test case pass\n", 1)
		} else {
			t.Errorf("Stack Fill should be 1 but is : %d\n", stack.Fill())
		}
	}
}

/**
  Test case to check the stack re initialization case.
*/
func TestReinstantiationCase(t *testing.T) {
	stack := ArrayStack{}
	err := stack.NewArrayStack(32)

	err = stack.NewArrayStack(16)
	if err == nil {
		t.Errorf("There must be an error when instantiating already existing stack use Reinitialize method instead")
	} else {
		t.Logf("Error while reinstantiating an array | Test PASS")
	}
}

func TestReinstantiationofStackArray(t *testing.T) {
	stack := ArrayStack{}
	err := stack.NewArrayStack(32)

	err = stack.Reinstantiate(16)
	if err != nil {
		t.Errorf("The reinitialization meth must allow clear and intatiate the new stack | Test FAIL")
	} else {
		t.Logf("No error while reinstantiating an stack | Test PASS")
	}
}
