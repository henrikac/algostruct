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

import "testing"

func TestNew(t *testing.T) {
	s := New()
	if s == nil {
		t.Error("The stack was not properly initialized")
	}
}

func TestNewTop(t *testing.T) {
	s := New()
	_, err := s.Peek()
	if err == nil {
		t.Error("Should not be able to peek an empty Stack")
	}
}

func TestNewLen(t *testing.T) {
	s := New()
	if s.Len() != 0 {
		t.Error("The stack was not properly initialized")
	}
}

type pushTest struct {
	in		interface{}
	length	int
}

func TestPush(t *testing.T) {
	items := []pushTest{
		{123, 1},
		{14, 2},
		{1, 3},
		{423413, 4},
	}
	s := New()
	for _, item := range items {
		s.Push(item.in)
		if s.Len() != item.length {
			t.Errorf("Expected stack length: %d\nActual stack length: %d", item.length, s.Len())
		}
		peek, _ := s.Peek()
		if peek != item.in {
			t.Errorf("Expected pushed item: %d\nActual pushed item: %d", peek, item.in)
		}
	}
}

func TestPop(t *testing.T) {
	s := New()
	expected := 7
	s.Push("Alice")
	s.Push(expected)

	item, _ := s.Pop()
	if item != expected {
		t.Errorf("Expected popped item to be %d", expected)
	}
	if s.Len() != 1 {
		t.Errorf("Expected stack length: %d", 1)
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := New()
	_, err := s.Pop()
	if err == nil {
		t.Error("Should not be able to pop items of an empty Stack")
	}
}
