package go_goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T)  {
	totalCpu := runtime.NumCPU()
	fmt.Println("CPU", totalCpu)

	//runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Goroutines", totalGoroutines)
}
