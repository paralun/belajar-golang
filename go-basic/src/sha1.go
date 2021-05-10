package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	text := "this is secret"
	sha := sha1.New()
	sha.Write([]byte(text))
	encrypted := sha.Sum(nil)
	encryptedString := fmt.Sprintf("%x", encrypted)
	fmt.Println(encryptedString)
}
