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

import "sort"

// Sorts a collection using bubble sort
func Sort(collection sort.Interface) {
	n := collection.Len()

	for n >= 1 {
		counter := 0
		for i := 1; i < n; i++ {
			if collection.Less(i, i - 1) {
				collection.Swap(i, i - 1)
				counter = i
			}
		}
		n = counter
	}
}

// Sorts a slice of integers using bubble sort
func Ints(ints []int) {
	Sort(sort.IntSlice(ints))
}

// Sorts a slice of strings using bubble sort
func Strings(strs []string) {
	Sort(sort.StringSlice(strs))
}

// Sorts a slice of float64 using bubble sort
func Floats(floats []float64) {
	Sort(sort.Float64Slice(floats))
}
