package main

import "fmt"

func sayYourName(firstNameVal string, lastNameVal string) (firstName string, lastName string) {
	firstName = firstNameVal
	lastName = lastNameVal

	return
}

func main() {
	firstName, lastName := sayYourName("Nabil", "Wafi")

	fmt.Println(firstName, lastName)
}