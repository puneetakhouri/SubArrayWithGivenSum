package main

import "fmt"

func main() {

	//var arr = []int{1, 2, 3, 7, 5}
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumToFind := 15

	found, startIndex, endIndex := FindSubArrayWithSum(arr, sumToFind)
	if found {
		fmt.Printf("The indexes are %d, %d\n", startIndex+1, endIndex+1)
	} else {
		fmt.Println("Could not find the sum")
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
		if j == len(arr) {
			return false, i, j
		}
	}
	return false, i, j
}
