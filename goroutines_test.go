package mylearngolanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func RunHelloworld(number int) {
	fmt.Println("hello world", number)
}

func TestCreateGoroutines(t *testing.T) {
	go RunHelloworld(1)
	fmt.Println("aww")

	time.Sleep(1 * time.Second)
}

func TestCountGoroutines(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go RunHelloworld(i)

	}
	time.Sleep(5 * time.Second)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "maher "
		fmt.Println("send data success to channel")
	}()

	data := <-channel
	go fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// channel as parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(200 * time.Millisecond)
	channel <- "maher zaenudin mukti umar"
}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

// channel for send data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "maher zaenuidin"
}

// channel for get data from func only in
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestOnlyInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "maiing"
		channel <- "only"
		channel <- "maher"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("selesai")
}

// how to use range to read data from a channel and close to shut it down
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
	}()

	for data := range channel {
		fmt.Println("menerima data", data)
	}
}

// sometimes there are cases where we create multiple channels and run go routines
// then we want to get data from some of these channels
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	count := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1 ", data)
			count++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2 ", data)
			count++
		//default select
		//waiting data from these channels or do something if data haven't entered the channel yet
		default:
			fmt.Println("Waiting data")
		}
		if count == 2 {
			break
		}
	}
}
