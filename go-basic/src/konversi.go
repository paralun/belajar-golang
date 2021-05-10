package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string -> int
	dataInt, _ := strconv.Atoi("200")
	fmt.Println(dataInt)

	// int -> string
	dataStr := strconv.Itoa(500)
	fmt.Println(dataStr)

	// base 10 -> numerik, 2 -> biner, 16 -> Hexa -> 8 -> Oktal
	// string -> int numerik (base 10) tipe data int64
	dataInt64, _ := strconv.ParseInt("164", 10, 64)
	fmt.Println(dataInt64)

	//string biner -> int
	dataInt8, _ := strconv.ParseInt("1010", 2, 8)
	fmt.Println(dataInt8)

	// string -> bool
	dataBool, _ := strconv.ParseBool("true")
	fmt.Println(dataBool)

	// base 10 -> numerik, 2 -> biner, 16 -> Hexa -> 8 -> Oktal
	// int -> hexa
	dataHexa := strconv.FormatInt(1000000, 16)
	fmt.Println(dataHexa)

	// cast string -> byte
	dataByte := []byte("datanya byte")
	fmt.Println(dataByte)

	// byte -> string
	dataStrByte := string([]byte{100, 97, 116, 97, 110, 121, 97, 32, 98, 121, 116, 101})
	fmt.Println(dataStrByte)
}
