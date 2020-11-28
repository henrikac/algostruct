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

package insertion

import (
	"sort"
	"testing"
)

func TestSortingInts(t *testing.T) {
	expected := sort.IntSlice{-1, 1, 2, 5, 5, 6, 8, 10}
	actual := sort.IntSlice{5, 2, 5, 6, 1, 10, -1, 8}
	Sort(actual)
	for i := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Expected %d\nActual %d", expected[i], actual[i])
		}
	}
}

func TestSortingStrings(t *testing.T) {
	expected := sort.StringSlice{"Alice", "Bob", "Johnny", "alice", "julian"}
	actual := sort.StringSlice{"Johnny", "julian", "alice", "Bob", "Alice"}
	Sort(actual)
	for i := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Expected %s\nActual %s", expected[i], actual[i])
		}
	}
}