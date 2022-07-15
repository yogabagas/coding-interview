package main

import (
	"fmt"
)

var (
	nums = []int{1, 2, 3, 1}
)

func main() {
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	return recursiveFunc(0, 0, nums)
}

func recursiveFunc(now int, next int, nums []int) bool {

	if now >= len(nums)-1 {
		return false
	}

	if now == next {
		return recursiveFunc(now, next+1, nums)
	}

	if nums[now] == nums[next] {
		return true
	}

	if next >= len(nums)-1 {
		return recursiveFunc(now+1, 0, nums)
	}

	return recursiveFunc(now, next+1, nums)
}

func withHash(limit int, num int) int {
	var hash int
	return (hash + num) * num % limit
}
