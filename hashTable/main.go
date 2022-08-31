package main

import "log"

type KeyVal struct {
	Key string
	Val interface{}
}

func main() {

	kv := KeyVal{
		Key: "1232",
		Val: "apple",
	}
	limit := 50

	st := hashTable(kv, limit)

	hs := setHashTable(st, kv.Val, limit)

	gh := getHashTable(st, hs)
	log.Println(gh, hs)

}

func hashTable(kv KeyVal, limit int) int {

	var hash int
	for i := 0; i < len(kv.Key)-1; i++ {
		hash = (hash + int([]byte(kv.Key)[i])) * i % limit
	}

	return hash

}

func setHashTable(key int, value interface{}, limit int) []interface{} {
	newArrays := make([]interface{}, limit)
	newArrays[key] = value

	return newArrays
}

func getHashTable(key int, list []interface{}) interface{} {
	return list[key]
}
