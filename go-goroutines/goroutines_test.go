package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func PrintHello()  {
	fmt.Println("Goroutines Run")
}

func TestGoroutines(t *testing.T)  {
	go PrintHello()
	fmt.Println("Selesai")
	time.Sleep(2 * time.Second)
}

func DisplayNumber(number int)  {
	fmt.Println("Number", number)
}

func TestManyGoroutines(t *testing.T)  {
	for i := 0; i <= 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}