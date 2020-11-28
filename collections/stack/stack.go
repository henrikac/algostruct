/* 
* Copyright (C) 2020 Henrik A. Christensen
* 
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
* 
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
* 
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package stack

import "errors"

type Stack struct {
	top	*node
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

// Returns the value of the node at the top of the Stack without removing it
// if the Stack is not empty; otherwise, returns an error
func (s *Stack) Peek() (interface{}, error) {
	if s.length < 1 {
		err := errors.New("Cannot peek an empty Stack")
		return nil, err
	}
	return s.top.value, nil
}

// Returns the number of nodes in the Stack
func (s *Stack) Len() int {
	return s.length
}

// Inserts a node at the top of the Stack
func (s *Stack) Push(item interface{}) {
	n := &node{item, s.top}
	s.top = n
	s.length++
}

// Removes and returns the value of the node at the top of the Stack
// if the Stack is not empty; otherwise, returns an error
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
