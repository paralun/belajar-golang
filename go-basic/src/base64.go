package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Ini bukan data"
	encode := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encode)

	decode, _ := base64.StdEncoding.DecodeString(encode)
	fmt.Println(string(decode))

	data2 := "https://kalipare.com/"
	encodeURL := base64.URLEncoding.EncodeToString([]byte(data2))
	fmt.Println(encodeURL)

	decodeURL, _ := base64.URLEncoding.DecodeString(encodeURL)
	fmt.Println(string(decodeURL))
}
