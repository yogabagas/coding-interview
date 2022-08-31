package main

import "log"

var (
	aNums = []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	bNums = []int{2, 1, 1, 2, 3, 5, 1, 2, 4}
	cNums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func main() {
	t := containFirstChar(aNums)
	log.Println(t)

	t1 := containFirstChar(bNums)
	log.Println(t1)

	t2 := containFirstChar(cNums)
	log.Println(t2)
}

func containFirstChar(arr []int) int {

	mapNum := make(map[int]bool)

	for i := 0; i < len(arr)-1; i++ {
		if _, ok := mapNum[arr[i]]; !ok {
			mapNum[arr[i]] = true
		} else {
			return arr[i]
		}
	}
	return 0
}
