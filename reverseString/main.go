package main

import "log"

var (
	name = "Yoga Bagasakthi"
)

func main() {

	log.Println(reverseString(name))

}

func reverseString(s string) string {

	var newString string

	for i := len(s) - 1; i >= 0; i-- {
		newString = newString + string(s[i])
	}

	return newString

}
