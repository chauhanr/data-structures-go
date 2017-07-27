package stacks

import (
	"fmt"
	"testing"
)

var testParamthesisSet = []struct {
	paranthesis   []string
	errorPosition []int
}{
	{[]string{"(", "(", "(", ")", ")", "(", ")", ")"}, nil},
	{[]string{")", "(", ")", "("}, []int{0, 3}},
	{[]string{"(", ")", "("}, []int{2}},
}

func TestParanthesisProblem(t *testing.T) {

	for _, testSet := range testParamthesisSet {
		errPositions := findErrorPositionForParathesis(testSet.paranthesis, "(", ")")
		fmt.Printf("%v\n", errPositions)
		errPosExpected := testSet.errorPosition
		if errPosExpected == nil && errPositions != nil {
			// do nothing
		} else if (errPosExpected == nil && errPositions != nil) || (errPosExpected != nil && errPositions == nil) {
			t.Errorf("For paranthesis %v Error positions mismatch expected: %v, but got: %v ", testSet.paranthesis, errPosExpected, errPositions)
		}
		if errPosExpected != nil && errPositions != nil {
			if len(errPosExpected) != len(errPositions) {
				t.Errorf("For paranthesis %v Error positions mismatch expected: %v, but got: %v ", testSet.paranthesis, errPosExpected, errPositions)
			} else {
				for ix, pos := range errPosExpected {
					if pos != errPositions[ix] {
						t.Errorf("For paranthesis %v expected Error index value %d, but got %d", testSet.paranthesis, pos, errPositions[ix])
					}
				}
			}
		}
	}
}

func findErrorPositionForParathesis(paranthensisCol []string, openBrace string, closeBrace string) []int {
	stack := ArrayStack{}
	if paranthensisCol == nil {
		return nil
	}
	errPos := []int{}
	stack.NewArrayStack(len(paranthensisCol))

	for index, brace := range paranthensisCol {
		if brace == openBrace {
			stack.Push(index)
		} else {
			// this means this is closing brace
			_, err := stack.Pop()
			if err != nil {
				// error means that the stack is empty
				errPos = append(errPos, index)
			}
		}
	}
	idx, popErr := stack.Pop()
	for popErr == nil {
		errPos = append(errPos, idx)
		idx, popErr = stack.Pop()
	}
	return errPos
}
