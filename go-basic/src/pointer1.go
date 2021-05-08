package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("Value A", numberA)
	fmt.Println("Address A", &numberA)
	fmt.Println("Value B", *numberB)
	fmt.Println("Address B", numberB)

	numberA = 10
	fmt.Println("Setelah melakukan perubahan")
	fmt.Println("Value A", numberA)
	fmt.Println("Address A", &numberA)
	fmt.Println("Value B", *numberB)
	fmt.Println("Address B", numberB)
}

