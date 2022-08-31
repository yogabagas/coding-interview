package main

import "fmt"

func validate(s string) bool {

	if string(s[0]) != "{" {
		return false
	}

	var foundStatment int
	for _, v := range s {

		if string(v) == "{" {
			foundStatment++
		}
		if string(v) == "}" {
			foundStatment--
		}
	}

	if foundStatment == 0 {
		return true
	}

	return false
}

func main() {
	v := validate("{{}}{}{}")
	fmt.Println(v)
}
