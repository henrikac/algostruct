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

package queue

import "errors"

type Queue struct {
	head	*node
	tail	*node
	length	int
}

type node struct {
	value	interface{}
	next	*node
}

// Initializes a new Queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Returns the length of the Queue
func (q *Queue) Len() int {
	return q.length
}

// Returns the value of the node at the beginning of the Queue without removing it
// if the Queue is not empty; otherwise, returns an error
func (q *Queue) Peek() (interface{}, error) {
	if q.length < 1 {
		err := errors.New("Cannot peek an empty Queue")
		return nil, err
	}
	return q.head.value, nil
}

// Adds a node to the end of the Queue
func (q *Queue) Enqueue(item interface{}) {
	n := &node{item, nil}
	if q.length == 0 {
		q.head = n
	} else {
		q.tail.next = n
	}
	q.tail = n
	q.length++
}

// Removes and returns the value of the node at the beginning of the Queue
// if the Queue is not empty; otherwise, returns an error
func (q *Queue) Dequeue() (interface{}, error) {
	if q.length < 1 {
		err := errors.New("Cannot peek an empty Queue")
		return nil, err
	}
	item := q.head
	if q.length == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
		item.next = nil
	}
	q.length--
	return item.value, nil
}
