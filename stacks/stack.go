package stacks

import "errors"

// ArrayStack is an implementation of a stack using array of integers here.
type ArrayStack struct {
	stack  []int
	fill   int
	length int
}

// NewArrayStack is a constructor that will help make the Stack.
func (s *ArrayStack) NewArrayStack(n int) error {
	if n <= 0 {
		return errors.New("Stack cannot be empty or of negative length")
	}
	if s.length > 0 {
		// this means that the array is already initialized and this method should not allow reinitialization.
		return errors.New("Cannot reinitialize an already initialized stack")
	}
	s.stack = make([]int, n)
	s.length = n
	return nil
}

// Reinstantiate method will clear the already instanted stack and then starts affresh.
func (s *ArrayStack) Reinstantiate(n int) error {
	s.Clear()
	s.length = 0
	err := s.NewArrayStack(n)
	return err
}

// Size gives a length fo the stack that has been instantiated.
func (s *ArrayStack) Size() int {
	return s.length
}

// Fill method gives the number of elements that are currently in the stakc.
func (s *ArrayStack) Fill() int {
	return s.fill
}

// Push method is the one that adds element to the stack.
func (s *ArrayStack) Push(element int) error {
	if s.length == 0 {
		return errors.New("cannot push elements to an uninstantiated stack | use NewArrayStack(n int) method to instantiate")
	}
	if s.fill == s.length {
		return errors.New("Stack is full cannot take any more elements")
	}
	s.stack[s.fill] = element
	s.fill++
	return nil
}

// Pop element method will pop an element from the stack.
func (s *ArrayStack) Pop() (int, error) {
	if s.length == 0 {
		return 0, errors.New("Cannot remove elements from an empty stack")
	}
	s.fill--
	value := s.stack[s.fill]
	return value, nil
}

// Clear all the elements in the stack
func (s *ArrayStack) Clear() {
	for i := 0; i < s.fill; i++ {
		s.stack[i] = 0
	}
	s.fill = 0
}
