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
