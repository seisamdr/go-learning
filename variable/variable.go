package main

import "fmt"

func main() {
	var lastName string  = "Doe"
	var firstName string = "John"

	lastName = "Bogie"

	// var bilBulat uint8 = 255
	// var bilDeci = 2.5

	// variabel pointer
	var numA int = 10
	var numB *int = &numA
	*numB = 8
	fmt.Print("Hello ", firstName, lastName, " ", numA, "!\n")
}