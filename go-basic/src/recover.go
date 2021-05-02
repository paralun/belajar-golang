package main

import "fmt"

// Recover digunakan untuk menangkap data Panic, apabila program terjadi error aplikasi akan tetap berjalan

func endApp3()  {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error:", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp3(error bool)  {
	defer endApp3()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp3(true)
}
