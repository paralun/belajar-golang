package main

import "fmt"

type HasName interface {
	GetName() string
}

func Hai(hasName HasName)  {
	fmt.Println("Hallo", hasName.GetName())
}

type Orang struct {
	Nama string
}

func (orang Orang) GetName() string {
	return orang.Nama
}

func main() {
	var orang1 Orang
	orang1.Nama = "Bambang"

	Hai(orang1)
}
