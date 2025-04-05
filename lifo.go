package jsonql

import (
	"fmt"
)

// Lifo - last in first out stack
type Lifo struct {
	Top  *Element
	Size int
}

// Element - an item in the stack
type Element struct {
	Value interface{}
	Next  *Element
}

// Stack - a set of functions of the Stack
type Stack interface {
	Len() int
	Push(value interface{})
	Pop() (value interface{})
	Peep() (value interface{})
	Print()
}

// Len - gets the length of the stack.
func (s *Lifo) Len() int {
	return s.Size
}

// Push - pushes the value into the stack.
func (s *Lifo) Push(value interface{}) {
	s.Top = &Element{value, s.Top}
	s.Size++
}

// Pop - pops the last value out of the stack.
func (s *Lifo) Pop() (value interface{}) {
	if s.Size > 0 {
		value, s.Top = s.Top.Value, s.Top.Next
		s.Size--
		return
	}
	return nil
}

// Peep - gets the last value in the stack without popping it out.
func (s *Lifo) Peep() (value interface{}) {
	if s.Size > 0 {
		value = s.Top.Value
		return
	}
	return nil
}

// Print - shows what's in the stack.
func (s *Lifo) Print() {
	tmp := s.Top
	for i := 0; i < s.Len(); i++ {
		fmt.Print(tmp.Value, ", ")
		tmp = tmp.Next
	}
	fmt.Println()
}
