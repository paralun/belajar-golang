package main

import "fmt"

// bisa menerima type data apapun
func Ups(i int) interface{}  {
	if i == 1 {
		return 1
	}else if i == 2 {
		return true
	}else {
		return "String"
	}
}

func main() {
	data := Ups(2)
	fmt.Println(data)
}
