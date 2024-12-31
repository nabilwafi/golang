package main_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Muhammad Nabil Wafi"
		fmt.Println("Successfully Sent Message From Goroutines")
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestChannelWithFunc(t *testing.T) {
	channel := make(chan string)

	go sendData(channel)

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func sendData(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Nabil Wafi"
}
func TestOnlyInOut(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Nabil Wafi"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	go func() {
		channel <- "Muhammad"
		channel <- "Nabil"
		channel <- "Wafi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go sendData(channel1)
	go sendData(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari case-1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari case-2:", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

	fmt.Println("Selesai")
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go sendData(channel1)
	go sendData(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari case-1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari case-2:", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}

	fmt.Println("Selesai")
}