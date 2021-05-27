package go_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// chanel sebagai alat komunikasi synchronous (blocking)
// channel seperti mekanisme async await

func TestCreateChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Data Channel"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep(4 * time.Second)
}

func ChannelAsParameter(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Data Channel"
}

func TestChannelAsParameter(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go ChannelAsParameter(channel)

	data := <- channel
	fmt.Println(data)
	time.Sleep(4 * time.Second)
}

func ChannelOnlyIn(channel chan<- string)  {
	time.Sleep(2 * time.Second)
	channel <- "Data Channel"
}

func ChannelOnlyOut(channel <-chan string)  {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go ChannelOnlyIn(channel)
	go ChannelOnlyOut(channel)

	time.Sleep(4 * time.Second)
}

func TestBufferedChannel(t *testing.T)  {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "Data 1"
	channel <- "Data 2"

	fmt.Println(<- channel)
	fmt.Println(<- channel)

	time.Sleep(4 * time.Second)
}

func TestRangeChnnel(t *testing.T)  {
	channel := make(chan string)

	go func() {
		for i := 0; i<= 10; i++ {
			channel <- "Data " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}
