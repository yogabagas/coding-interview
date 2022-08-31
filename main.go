package main

import "log"

var (
	a = []int{1, 2, 3}
)

func main() {

	// s := increaseNumber(5, 8, 7)
	// log.Println(s)

	log.Println(Solution(a))

	// s := Hash("6F57F864D6C244AAA6D4F2ABF0E7B923")
	// vs := hex.EncodeToString(s[:])

	// log.Println(vs)

	// t:= hash58("6F57F864D6C244AAA6D4F2ABF0E7B923")

}

func increaseNumber(firstNum, secondNum, length int) []int {

	resp := make([]int, length)
	x := secondNum - firstNum

	resp[0], resp[1] = firstNum, secondNum

	for i := 1; i < length-1; i++ {
		resp[i+1] = resp[i] + x
	}

	return resp
}

func Solution(a []int) int {

	s := make(map[int]bool, len(a))

	for i := 0; i < len(a); i++ {
		s[a[i]] = true
	}

	log.Println(s, len(s))

	for i := 1; i < len(s); i++ {
		if _, ok := s[i]; !ok {
			return i
		}
	}

	return 0
}

// func Hash(v interface{}) [32]byte {
// 	val := fmt.Sprintf("%v", v)
// 	return sha256.Sum256([]byte(val))
// }

// func hash58(str string) string {
// 	h := sha256.Sum256([]byte(str))
// 	return b58.Encode(h[:])
// }
