package main

import "fmt"

// Function Defer akan selalu dieksekusi walapun terjadi error

func logging()  {
	fmt.Println("Melakukan Logging")
}

func runApp(value int)  {
	defer logging() // akan selalu di panggil di akhir
	fmt.Println("Run Appliction")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApp(0) // value 0 biar error
}
