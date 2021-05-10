package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Data1")
	data.PushBack("Data2")
	data.PushBack("Data3")
	data.PushBack("Data4")

	// menambahkan data dari depan
	data.PushFront("Data0")

	// mengambil data pertama
	fmt.Println(data.Front().Value)
	// mengambil data terakir
	fmt.Println(data.Back().Value)
	fmt.Println("")

	// data dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("")
	// data dari depan ke belakang
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

}
