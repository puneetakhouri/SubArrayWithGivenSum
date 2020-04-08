package main

import "fmt"

func main() {

	//var arr = []int{1, 2, 3, 7, 5}
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sumToFind := 24
	var i, j int = 0, 0
	for i <= j {
		sum := 0
		for k := i; k <= j; k++ {
			sum = sum + arr[k]
		}
		if sum == sumToFind {
			fmt.Printf("The indexes are %d, %d\n", i+1, j+1)
			break
		} else if sum < sumToFind {
			j = j + 1
		} else {
			i = i + 1
		}
	}
}
