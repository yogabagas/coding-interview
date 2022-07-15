package main

import (
	"fmt"
)

var (
 arr = []int{2,3,4,5,1}
 target = 9
)

func main() {
	fmt.Println(twoSum(arr, target))

}

func twoSum(nums []int, target int) []int {
	return recursiveSum(0, 0, nums, target)
}

func recursiveSum(now, next int, nums []int, target int) []int {
	newArrays := make([]int, 2)

	if now >= len(nums)-1 {
		return newArrays
	}

	if next <= now {
		return recursiveSum(now, next+1, nums, target)
	}

	if nums[now]+nums[next] == target {
		newArrays[0], newArrays[1] = nums[now], nums[next]
		return newArrays
	}

	if next >= len(nums)-1 {
		return recursiveSum(now+1, 0, nums, target)
	}

	return recursiveSum(now, next+1, nums, target)

}
