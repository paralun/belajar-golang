package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToLower("MENJADI KECIL"))
	fmt.Println(strings.ToUpper("menjadi besar"))
	fmt.Println(strings.Contains("Aisyah Saraswati", "Saraswati"))
	fmt.Println(strings.Trim("  Test Saja   ", " "))
	fmt.Println(strings.ReplaceAll("coba urang replace", "urang", "kita"))
	fmt.Println(strings.Split("data/rahasia", "/"))
}
