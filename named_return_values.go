package main

import "fmt"

func getCompleteName() (firstName , middleName , lastName string) {
	firstName = "Leon"
	middleName = "Siahaan"
	// lastName = "Sitinjak"
	
	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}