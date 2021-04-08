package main

import "fmt"

func main() {

	var a = 100
	var b = 10

	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var h = 20
	h += 10
	fmt.Println(h)

	h++
	fmt.Println(h)

}