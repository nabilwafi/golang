package main_test

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	counter := 0

	for i := 0; i < 1000; i++ {
		go func () {
			for i := 0; i < 100; i++ {
				counter++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", counter)
}