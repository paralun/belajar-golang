package main

import "fmt"

func main() {
	var nama string
	nama = "Siapa Yaaa..."
	fmt.Println(nama)

	var alamat = "Dimana Yaa.."
	fmt.Println(alamat)

	var umur = 1000
	fmt.Println(umur)

	var nomer int8 = 12
	fmt.Println(nomer)

	// var bentuknya optional
	nama2 := "Siapa lagi yaa.."
	fmt.Println(nama2)

	// Variable Multiple
	var (
		firstName = "Nama Pertama"
		lastName = "Nama Akhir"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)


	// Variable Constant
	const var1  = "Variable1"
	const var2  = 2000

	const (
		var3 = "Variable3"
		var4 = 9090
	)

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
}
