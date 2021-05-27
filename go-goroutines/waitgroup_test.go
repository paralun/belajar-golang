package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchrounous(group *sync.WaitGroup)  {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T)  {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++  {
		go RunAsynchrounous(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}
