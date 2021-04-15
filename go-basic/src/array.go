package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "Nisa"
	names[1] = "Sofia"
	names[2] = "Agus"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int {
		90,
		40,
		20,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
