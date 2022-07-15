package main

import (
	"fmt"
	"log"
)

var (
	nums = []int{0, 1, 0, 3, 12}
)

func main() {
	fmt.Println(recursiveFunc(0, 0, nums))
}

func moveZereos(n []int) {
	for i := 0; i < len(n)-1; i++ {
		for j := 0; j <= len(n)-1; j++ {

			if j == i {
				continue
			}

			if n[i] == 0 && j > i {
				n[j], n[i] = n[i], n[j]
			}
		}
	}

	log.Println(n)

}

func moveZeroesWithSort(n []int) {

	for i := 0; i < len(n)-1; i++ {
		for j := 0; j <= len(n)-1; j++ {

			if j == i {
				continue
			}

			if ((n[j] < n[i] && n[j] != 0) || n[i] == 0) && j > i {
				n[j], n[i] = n[i], n[j]
			}
		}

	}

	log.Println(n)

}

func recursiveFunc(now int, next int, n []int) []int {

	if now >= len(n)-1 {
		return n
	}

	if next > len(n)-1 {
		return recursiveFunc(now+1, 0, n)
	}

	if now == next {
		return recursiveFunc(now, next+1, n)
	}

	if n[now] == 0 && next > now {
		n[next], n[now] = n[now], n[next]
	}

	return recursiveFunc(now, next+1, n)
}
