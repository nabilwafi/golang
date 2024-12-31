package main

import "fmt"

type Biodata struct {
	Name, Hobby string
}

func changeHobby(biodata *Biodata, value string) {
	biodata.Hobby = value
}

func main() {

	var user Biodata = Biodata{
		Name:  "Muhammad Nabil Wafi",
		Hobby: "Chess",
	}

	fmt.Println(user)
	changeHobby(&user, "Swimming")
	fmt.Println(user)
}