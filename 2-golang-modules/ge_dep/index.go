package main

import (
	go_say_hello "github.com/testingnabilwafi/go-say-hello/v2"
	"fmt"
)

func main() {
	val := go_say_hello.SayHello()

	fmt.Println(val)
}
