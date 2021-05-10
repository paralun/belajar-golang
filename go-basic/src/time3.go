package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	start := time.Now()
	time.Sleep(time.Second * 5)
	duration := time.Since(start)

	fmt.Println(duration.Seconds())
	fmt.Println(duration.Minutes())
}
