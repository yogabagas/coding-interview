package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	b58 "github.com/jbenet/go-base58"
)

func main() {

	s := Hash("50")
	vs := hex.EncodeToString(s[:])

	log.Println(vs)

}

func Hash(v interface{}) [32]byte {
	val := fmt.Sprintf("%v", v)
	return sha256.Sum256([]byte(val))
}

func hash58(str string) string {
	h := sha256.Sum256([]byte(str))
	return b58.Encode(h[:])
}
