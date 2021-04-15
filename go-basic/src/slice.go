package main

import "fmt"

func main() {
	// Tipe Data Slice adalah potongan dari data array
	// Slice mirip dengan array yg membedakan adalah ukuran slice bisa berubah
	// Slice memiliki 3 data yaitu pointer, length dan capacity

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// jika merubah data array, slice akan ikut berubah begitu juga sebaliknya
	months[4] = "Jupiter"
	fmt.Println(slice1)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[3:4]
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Liburrrr")
	fmt.Println(daySlice2)
	fmt.Println(days)

	// buat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "satu"
	newSlice[1] = "dua"
	fmt.Println(newSlice)

	// copy data slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// perbedaan deklalarsi array dengan slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
