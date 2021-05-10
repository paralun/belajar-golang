package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	fmt.Println(arg)

	host, _ := os.Hostname()
	fmt.Println(host)
}
