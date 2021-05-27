package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int)  {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestConditionSignal(t *testing.T)  {
	for i := 1; i <= 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	group.Wait()
	fmt.Println("Selesai")
}

func TestConditionBroadcast(t *testing.T)  {
	for i := 1; i <= 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
	fmt.Println("Selesai")
}
