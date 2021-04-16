package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	}

	hari := "Minggu"
	switch hari {
	case "Sabtu", "Minggu":
		fmt.Println("Hari Libur")
	default:
		fmt.Println("Hari Kerja")
	}

	nama := "Bambang"
	switch jml := len(nama); jml > 6 {
	case true:
		fmt.Println("OK")
	case false:
		fmt.Println("NOT OK")
	}

	jam := 10
	switch {
	case jam < 12:
		fmt.Println("Sebelum tengah hari")
	default:
		fmt.Println("Setelah tengah hari")
	}

	siapaSaya := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("Saya Boolean")
		case int:
			fmt.Println("Saya Integer")
		default:
			fmt.Println("Tidak dikenal")
		}
	}

	siapaSaya(true)
	siapaSaya(20)
	siapaSaya("hey")
}
