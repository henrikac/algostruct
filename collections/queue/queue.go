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
