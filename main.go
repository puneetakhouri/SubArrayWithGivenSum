package main

import "fmt"

func main() {

	//var arr = []int{1, 2, 3, 7, 5}
	//var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var arr = []int{135, 101, 170, 125, 79, 159, 163, 65, 106, 146, 82, 28, 162, 92, 196, 143, 28, 37, 192, 5, 103, 154, 93, 183, 22, 117, 119, 96, 48, 127, 172, 139, 70, 113, 68, 100, 36, 95, 104, 12, 123, 134}
	//sumToFind := 468

	var arr = []int{10, 9, 8}
	sumToFind := 8

	found, startIndex, endIndex := FindSubArrayWithSum(arr, sumToFind)
	if found {
		fmt.Printf("The indexes are %d, %d\n", startIndex+1, endIndex+1)
	} else {
		fmt.Println("-1")
	}
}

//FindSubArrayWithSum method does the main calculation
func FindSubArrayWithSum(arr []int, sumToFind int) (bool, int, int) {
	var i, j int = 0, 0
	for i <= j {
		sum := 0
		for k := i; k <= j; k++ {
			sum = sum + arr[k]
		}
		if sum == sumToFind {
			return true, i, j
		} else if sum < sumToFind {
			j = j + 1
		} else {
			i = i + 1
		}
		if j < i && j < len(arr) {
			j = j + 1
		}
		if j == len(arr) {
			return false, i, j
		}
	}
	return false, i, j
}
