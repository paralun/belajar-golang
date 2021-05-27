package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup)  {
	defer group.Done()

	group.Add(1)

	data.Store(value, value)
}

func TestMap(t *testing.T)  {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++  {
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println("Key", key, "Value", value)
		return true
	})
}
