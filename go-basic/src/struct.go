package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

// struct method
func (customer Customer) getName()  {
	fmt.Println("Hallooo", customer.Name)
}

func main() {
	var cus1 Customer
	cus1.Name = "Anissa"
	cus1.Address = "Cianjur"
	cus1.Age = 25

	fmt.Println(cus1)
	fmt.Println(cus1.Name)
	fmt.Println(cus1.Address)
	fmt.Println(cus1.Age)

	cus2 := Customer{
		Name:    "Budi",
		Address: "Bandung",
		Age:     30,
	}
	fmt.Println(cus2)

	//harus sama urutannya
	cus3 := Customer{"Bella", "Karawang", 22}
	fmt.Println(cus3)

	// manggil struct method
	cus1.getName()
}
