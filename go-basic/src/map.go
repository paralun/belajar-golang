package main

import "fmt"

func main() {

	// map[type key]type value{data map}
	person := map[string]string{
		"nama" : "Bambang",
		"alamat" : "Bandung",
	}

	// menambah data
	person["title"] = "Pengangguran"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person["title"])

	// buat map kosongan
	//var data map[string] string = make(map[string] string)
	data := make(map[string]string)
	data["satu"] = "Satu"
	data["dua"] = "Dua"
	data["salah"] = "Dihapus"

	delete(data, "salah")

	fmt.Println(data)
	fmt.Println(len(data))
}
