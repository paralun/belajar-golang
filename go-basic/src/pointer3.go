package main

import "fmt"

type Address struct {
	nama string
	alamat string
	kota string
}

// pointer di function
func changeAlamat(address *Address)  {
	address.kota = "Jakarta"
}

// pointer di method struct
func (addr *Address) changeNama()  {
	addr.nama = "Mr/Mrs " + addr.nama
}

func main() {

	addr1 := Address{
		nama:   "Bambang",
		alamat: "Jl. Belok Sebelah",
		kota:   "bandung",
	}

	addr2 := &addr1
	var addr3 *Address = &addr1

	addr2.kota = "Cianjur"

	// buat pointer baru, data aslinya tidak berubah
	//addr2 = &Address{"Aisyah", "Jl. Mahmud", "Sukabumi"}

	// data aslinya akan ikut berubah
	*addr2 = Address{"Aisyah", "Jl. Mahmud", "Sukabumi"}

	fmt.Println(addr1)
	fmt.Println(addr2)
	fmt.Println(addr3)

	// pointer kosong
	var addr4 *Address = new(Address)
	fmt.Println(addr4)

	var alamat = Address{"Udin", "Jl. Bidakarya", "Medan"}
	changeAlamat(&alamat)
	fmt.Println(alamat)

	// pointer struct method
	var alamat2 = Address{"Jarwo", "Jl. Sayana", "Kalimantan"}
	alamat2.changeNama()
	fmt.Println(alamat2)
}
