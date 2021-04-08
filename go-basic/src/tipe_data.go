package main

import (
	"fmt"
	"strings"
)

func main() {
	// Number
	fmt.Println(1)
	fmt.Println(100)
	fmt.Println(10.6)

	// Boolean
	fmt.Println(true)
	fmt.Println(false)

	// String
	fmt.Println("Ini String")
	fmt.Println(len("ABC"))
	fmt.Println(strings.ToUpper("upper"))
	fmt.Println(strings.ToLower("LOWER"))

	// Konversi type data
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var nama = "Kusmambang"
	var by byte = nama[1]
	var kar string = string(by)

	fmt.Println(kar)
}
