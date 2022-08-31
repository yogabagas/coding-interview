package main

import (
	"log"
	"math"
)

var mapDices map[int]int

func main() {
	n := numberDices(6, []int{1, 4, 4})
	log.Println(n)
}

func numberDices(n int, dices []int) int {
	minSwap := math.MaxInt
	for i := 0; i <= 6; i++ {
		currentStep := 0

		for j := 0; j < n; j++ {
			currentStep += getStepChange(dices[j], i)
		}
		minSwap = int(math.Min(float64(minSwap), float64(currentStep)))
	}
	return minSwap
}

func getStepChange(a, b int) int {
	mapDices = map[int]int{
		0: 0,
		1: 6,
		2: 5,
		3: 4,
	}

	if a == b {
		return 0
	}

	if c, ok := mapDices[a]; ok {
		if c == b {
			return 2
		}
		return 1
	}
	return -1
}
