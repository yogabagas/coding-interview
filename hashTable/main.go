package main

import "log"

type KeyVal struct {
	Key string
	Val interface{}
}

func main() {

	hashTable(KeyVal{
		Key: "1232",
		Val: 1000,
	}, 50)

}

func hashTable(kv KeyVal, limit int) {

	var hash int
	for i := 0; i < len(kv.Key)-1; i++ {
		hash = (hash + int([]byte(kv.Key)[i])) * i % limit
	}

	log.Println(hash)

}

func setHashTable(key int, value interface{}, limit int) []interface{} {
	newArrays := make([]interface{}, limit)
	newArrays[key] = value
	return newArrays
}
