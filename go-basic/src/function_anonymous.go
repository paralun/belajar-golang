package main

import "fmt"

type Blacklist func(string) bool

func registerUser(nama string, blacklist Blacklist)  {
	if blacklist(nama) {
		fmt.Println("Anda di block", nama)
	}else {
		fmt.Println("Selamat..", nama)
	}
}

func main() {
	blacklist := func(nama string) bool {
		return nama == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("baba", blacklist)

	registerUser("root", func(nama string) bool {
		return nama == "root"
	})
}
