package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T)  {
	group := sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Data1")
	pool.Put("Data2")
	pool.Put("Data3")

	for i := 1; i <= 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Selesai")
}
