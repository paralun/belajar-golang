package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func SelectChannel(channel chan string)  {
	channel <- "Data Channel"
	time.Sleep(1 * time.Second)
}

func TestSelectChannel(t *testing.T)  {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go SelectChannel(channel1)
	go SelectChannel(channel2)

	select {
	case data := <- channel1:
		fmt.Println("Channel 1", data)
	case data := <- channel2:
		fmt.Println("Channel 2", data)
	}

	select {
	case data := <- channel1:
		fmt.Println("Channel 1", data)
	case data := <- channel2:
		fmt.Println("Channel 2", data)
	}

	time.Sleep(3 * time.Second)
}

func TestSelectChannel2(t *testing.T)  {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go SelectChannel(channel1)
	go SelectChannel(channel2)

	counter := 0
	for {
		select {
		case data := <- channel1:
			fmt.Println("Channel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("Channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T)  {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go SelectChannel(channel1)
	go SelectChannel(channel2)

	counter := 0
	for {
		select {
		case data := <- channel1:
			fmt.Println("Channel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}