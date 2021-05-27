package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func PrintLoop(msg string)  {
	for i := 0; i <= 100; i++ {
		fmt.Println(msg, i)
	}
}

func TestGo(t *testing.T)  {
	go PrintLoop("Message A")
	go PrintLoop("Message B")

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}