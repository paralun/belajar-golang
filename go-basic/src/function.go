package main

import (
	"fmt"
	"strings"
)

func hayBroo()  {
	fmt.Println("Hallo Brooo...")
}

func hayBray(nama string)  {
	fmt.Println("Hay Brayy", nama)
}

func getHallo(nama string)string  {
	return "Hallo " + nama
}

// Multiple return value
func getFullName()(string,string)  {
	return "Pertama", "Kedua"
}

// Return named value
func getPerson()(nama string, alamat string, umur int)  {
	nama = "Juwita"
	alamat = "Cikota"
	umur = 23

	return
}

// Variadic Function
func sum(numbers ...int) int  {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

//Function Value
func getGoodBye(nama string)string  {
	return "Good Bye " + nama
}

// Function as Parameter
func sayWithFilter(nama string, filter func(string)string)  {
	namaFilter := filter(nama)
	fmt.Println(namaFilter)
}

func upFilter(nama string)string  {
	return strings.ToUpper(nama)
}

func main() {
	hayBroo()
	hayBray("Bambanx")

	result := getHallo("Jubaedah")
	fmt.Println(result)

	pertama, kedua := getFullName()
	fmt.Println(pertama, kedua)

	// mengabaikan return value
	_, b := getFullName()
	fmt.Println(b)

	nama, alamat, umur := getPerson()
	fmt.Println(nama, alamat, umur)

	total := sum(10, 20, 30, 40, 50)
	fmt.Println(total)

	slice2 := []int{10, 30, 50}
	total2 := sum(slice2...)
	fmt.Println(total2)

	// Function dimasukan kedalam variable
	goodbye := getGoodBye
	fmt.Println(goodbye("Bambanx"))

	// Function as Parameter
	sayWithFilter("ujang", upFilter)
}
