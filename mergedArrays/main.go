package main

import (
	"fmt"
)

var (
	n1 = []int{2, 4, 3, 5, 1}
	n2 = []int{3, 6, 7, 9, 1}
)

func main() {
	fmt.Println(mergedArrays(n1, n2))
}

func mergedArrays(n1 []int, n2 []int) []int {

	n1 = recursiveSort(0, 0, n1)
	n2 = recursiveSort(0, 0, n2)

	for _, w := range n2 {
		n1 = append(n1, w)
	}

	n1 = recursiveSort(0, 0, n1)

	return n1
}

func recursiveSort(now int, next int, n []int) []int {

	if now >= len(n)-1 {
		return n
	}

	if now == next {
		return recursiveSort(now, next+1, n)
	}

	if n[now] > n[next] && now < next {
		n[now], n[next] = n[next], n[now]
	}

	if n[now] < n[next] && now > next {
		n[next], n[now] = n[now], n[next]
	}

	if next >= len(n)-1 {
		return recursiveSort(now+1, 0, n)
	}

	return recursiveSort(now, next+1, n)

}
