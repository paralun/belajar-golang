package main

import "fmt"

func main() {

	nilai := 70

	if nilai >= 80 {
		fmt.Println("SANGAT BAIK")
	} else if nilai >= 60 {
		fmt.Println("BAIK")
	} else if nilai >= 40 {
		fmt.Println("CUKUP")
	} else {
		fmt.Println("GORENG")
	}

	// Short Statement
	no := "1234567891"
	if length := len(no); length > 12 {
		fmt.Println("No kepanjangan")
	}else {
		fmt.Println("Sippppp")
	}
}
