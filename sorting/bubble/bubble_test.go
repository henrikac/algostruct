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

package bubble

import "testing"

type testPerson struct {
	name	string
	age	int
}

type personSlice []testPerson

func (ps personSlice) Len() int {
	return len(ps)
}

func (ps personSlice) Less(i, j int) bool {
	if ps[i].age < ps[j].age {
		return true
	}
	if ps[i].age == ps[j].age {
		return ps[i].name < ps[j].name
	}
	return false
}

func (ps personSlice) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func TestSortingStructs(t *testing.T) {
	expected := personSlice{
		testPerson{"Johnny", 4},
		testPerson{"Bob", 12},
		testPerson{"Bob", 35},
		testPerson{"alice", 35},
		testPerson{"Alice", 89},
	}
	actual := personSlice{
		testPerson{"alice", 35},
		testPerson{"Bob", 35},
		testPerson{"Bob", 12},
		testPerson{"Alice", 89},
		testPerson{"Johnny", 4},
	}
	Sort(actual)
	for i := range actual {
		if expected[i].age != actual[i].age && expected[i].name != actual[i].name {
			t.Errorf("Expected: %s %d\nActual: %s %d",
				expected[i].name,
				expected[i].age,
				actual[i].name,
				actual[i].age,
			)
		}
	}
}

func TestSortingInts(t *testing.T) {
	expected := []int{-1, 1, 2, 5, 5, 6, 8, 10}
	actual := []int{5, 2, 5, 6, 1, 10, -1, 8}
	Ints(actual)
	for i := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Expected %d\nActual %d", expected[i], actual[i])
		}
	}
}

func TestSortingStrings(t *testing.T) {
	expected := []string{"Alice", "Bob", "Johnny", "alice", "julian"}
	actual := []string{"Johnny", "julian", "alice", "Bob", "Alice"}
	Strings(actual)
	for i := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Expected %s\nActual %s", expected[i], actual[i])
		}
	}
}

func TestSortingFloats(t *testing.T) {
	expected := []float64{0.0, 1.9, 3.7, 5.2, 5.2, 8.4}
	actual := []float64{5.2, 8.4, 5.2, 1.9, 0.0, 3.7}
	Floats(actual)
	for i := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Expected %f\nActual %f", expected[i], actual[i])
		}
	}
}
