package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	
	slice1 := []string{"Data1", "Data2", "Data3"}
	for j := 0; j < len(slice1); j++  {
		fmt.Println(slice1[j])
	}

	// menggunakan Range
	slice2 := []string{"Data1", "Data2", "Data3", "Data4"}
	for i, value := range slice2 {
		fmt.Println("Index", i, "=", value)
	}

	map1 := make(map[string]string)
	map1["satu"] = "Satu"
	map1["dua"] = "Dua"
	map1["tiga"] = "Tiga"

	for key, value := range map1 {
		fmt.Println(key, "=", value)
	}

	// Break
	for x := 1; x <= 10; x++  {
		if x == 5 {
			break
		}

		fmt.Println("Perulangan ke - ", x)
	}

	// Continue
	for z := 1; z <= 10; z++  {
		if z % 2 == 0 {
			continue
		}

		fmt.Println("Hasil ke - ", z)
	}
}
