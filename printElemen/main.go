package main

import "fmt"

func main() {
	printOutElem(20)
}

func printOutElem(n int) {

	for i := 1; i <= n; i++ {
		if i == 2 {
			fmt.Println("*")
		} else if i%5 == 0 {
			fmt.Println("*")
		} else if i%5 == 2 {
			fmt.Println("*")
		} else {
			fmt.Println(i)
		}
	}
}
