package main

import (
	"crypto/sha256"
	"fmt"
)

func GenerateEmojiHash(data string) [32]byte {
	sum := sha256.Sum256([]byte(data))
	fmt.Printf("%x\n", sum)
	return sum
}

func main() {
	s := GenerateEmojiHash("Hello this is Pratik")
	fmt.Println(s)
}
