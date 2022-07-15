package main

import (
	"fmt"
)

var (
	arr = []int{1, 2, 3, 4, 5, 6, 7}
	n   = 3
)

func main() {

	fmt.Println(rotateArray(n, arr))

}

func rotateArray(n int, arr []int) []int {
	newArrays := make([]int, len(arr))
	for i := 1; i <= n; i++ {

		lastIndex := arr[len(newArrays)-1]
		newArrays = make([]int, len(arr))

		for j := 0; j <= len(arr)-1; j++ {
			newArrays[0] = lastIndex
			if j > 0 {
				newArrays[j] = arr[j-1]
			}
		}
		arr = newArrays

	}
	return newArrays
}
