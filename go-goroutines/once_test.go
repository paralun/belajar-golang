package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce()  {
	counter++
}

func TestOnce(t *testing.T)  {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}