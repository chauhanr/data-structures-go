package stacks

import (
	"fmt"
)

type StackElement interface {
	Element() interface{}
	PrintElement()
}

type CharElement struct {
	element string
}

type IntElement struct {
	element int
}

type FloatElement struct {
	element float64
}

func (i IntElement) Element() interface{} {
	return i.element
}

func (i IntElement) PrintElement() {
	fmt.Printf("%d ", i.element)
}

func (c CharElement) Element() interface{} {
	return c.element
}

func (c CharElement) PrintElement() {
	fmt.Printf("%s ", c.element)
}
