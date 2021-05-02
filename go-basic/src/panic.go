package main

import "fmt"

// Func Panic digunakan untuk mengentikan program, biasanya ketika terjadi error

func endApp()  {
	fmt.Println("Aplikasi Selesai")
}

func runApp2(error bool)  {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp2(true)
}
