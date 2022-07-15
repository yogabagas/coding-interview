// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

var numbers1 = []int{1, 2, 3, 5, 9} // should return true
var numbers2 = []int{1, 2, 4, 5, 6} // should return false
var sum = 9

func main() {
	now := time.Now()
	fmt.Println(havingSum(0, 0, sum, numbers1))
	fmt.Println(havingSum(0, 0, sum, numbers2))

	fmt.Println(time.Since(now).Microseconds())
}

func havingSum(now int, next int, sum int, ns []int) bool {

	if now > len(ns)-1 {
		return false
	}

	if now == next {
		return havingSum(now, next+1, sum, ns)
	}

	if next > len(ns)-1 {
		return havingSum(now+1, 0, sum, ns)
	}

	if ns[now]+ns[next] == sum {
		return true
	}

	return havingSum(now, next+1, sum, ns)
}
