package main

import (
	"strconv"
	"testing"
)

func TestFindSubArrayWithSum(t *testing.T) {
	tables := []struct {
		arr       []int
		sumToFind int
		found     bool
	}{
		{[]int{1, 2, 3, 7, 5}, 12, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 15, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 99, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 59, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 252576}, 252576, true},
		{[]int{135, 101, 170, 125, 79, 159, 163, 65, 106, 146, 82, 28, 162, 92, 196, 143, 28, 37, 192, 5, 103, 154, 93, 183, 22, 117, 119, 96, 48, 127, 172, 139, 70, 113, 68, 100, 36, 95, 104, 12, 123, 134}, 468, true},
		{[]int{240, 201, 173}, 173, true},
	}

	for _, table := range tables {
		found, _, _ := FindSubArrayWithSum(table.arr, table.sumToFind)
		if found == table.found {
			t.Log("Success")
		} else {
			t.Error("Error, expected " + strconv.FormatBool(table.found) + ", got " + strconv.FormatBool(found))
		}
	}
}
