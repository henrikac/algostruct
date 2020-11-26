package stack

import "errors"

type Stack struct {
	top		*node
	length	int
}

type node struct {
	value	interface{}
	prev	*node
}

// Initializes a new Stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Returns the item at the top of the Stack without removing it
// Returns an error if the Stack is empty
func (s *Stack) Peek() (interface{}, error) {
	if s.length < 1 {
		err := errors.New("Cannot peek an empty Stack")
		return nil, err
	}
	return s.top.value, nil
}

// Returns the number of items in the Stack
func (s *Stack) Len() int {
	return s.length
}

// Inserts an item at the top of the Stack
func (s *Stack) Push(item interface{}) {
	n := &node{item, s.top}
	s.top = n
	s.length++
}

// Removes and returns the item at the top of the Stack
// Returns an error if the Stack is empty
func (s *Stack) Pop() (interface{}, error) {
	if s.length == 0 {
		err := errors.New("Cannot pop items of an empty Stack")
		return nil, err
	}
	item := s.top
	s.top = item.prev
	s.length--
	return item.value, nil
}
