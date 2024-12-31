package main_test

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(number int) {
	fmt.Println("Display:", number)
}

func TestDisplay(t *testing.T) {

	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}


