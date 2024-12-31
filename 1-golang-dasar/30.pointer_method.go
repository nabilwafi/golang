package main

import "fmt"

type User struct {
	Name string
}

func (user *User) Married() {
	user.Name = "Mr. " + user.Name
}

func main() {
	user := User{"Nabil"}

	user.Married()
	fmt.Println(user)
}