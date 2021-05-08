package main

import "fmt"

func change(original *int, value int)  {
	*original = value
}

func main() {
	number := 4
	fmt.Println("Before :", number)

	change(&number, 10)
	fmt.Println("After", number)
}
