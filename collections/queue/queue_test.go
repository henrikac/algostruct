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

import "testing"

func TestNew(t *testing.T) {
	q := New()
	if q == nil {
		t.Error("The queue was not properly initialized")
	}
}

func TestNewQueueLen(t *testing.T) {
	q := New()
	if q.length != 0 {
		t.Errorf("Expected a new Queue to have length %d but actual length is %d", 0, q.length)
	}
}

func TestPeekEmptyQueue(t *testing.T) {
	q := New()
	_, err := q.Peek()
	if err == nil {
		t.Error("Should not be able to peek an empty Queue")
	}
}

func TestEnqueue(t *testing.T) {
	q := New()
	q.Enqueue("Alice")
	q.Enqueue("Bob")
	if q.length != 2 {
		t.Errorf("Expected Queue to contain %d items but actual contains %d item(s)", 2, q.length)
	}
	head := q.head.value
	if head != "Alice" {
		t.Errorf("Expected %s\nActual %s", "Alice", head)
	}
	tail := q.tail.value
	if tail != "Bob" {
		t.Errorf("Expected %s\nActual %s", "Bob", tail)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := New()
	_, err := q.Dequeue()
	if err == nil {
		t.Error("Should not be able to dequeue an empty Queue")
	}
}

func TestDequeueOnQueueWithJustOneItem(t *testing.T) {
	q := New()
	q.Enqueue(123)
	item, _ := q.Dequeue()
	if item != 123 {
		t.Errorf("Expected %d\nActual %d", 123, item)
	}
	if q.length != 0 {
		t.Errorf("Queue should be empty")
	}
	if q.head != nil {
		t.Error("Head should be nil")
	}
	if q.tail != nil {
		t.Error("Tail should be nil")
	}
}

func TestDequeue(t *testing.T) {
	q := New()
	q.Enqueue("Bob")
	q.Enqueue("Alice")
	q.Enqueue("John")
	item, _ := q.Dequeue()
	if item != "Bob" {
		t.Errorf("Expected %s\nActual %s", "Bob", item)
	}
	if q.length != 2 {
		t.Errorf("Expected Queue to contain %d items but actual contains %d item(s)", 2, q.length)
	}
	head := q.head.value
	if head != "Alice" {
		t.Errorf("Expected %s\nActual %s", "Alice", head)
	}
	tail := q.tail.value
	if tail != "John" {
		t.Errorf("Expected %s\nActual %s", "John", tail)
	}
}
