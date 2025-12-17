package	main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Leon"
	// person["address"] = "Medan"
	// fmt.Println(person)

	person := map[string]string{
		"name" : "Leon",
		"address" : "Medan",
	}
	
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Leon Siahaan"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	
	fmt.Println(book)
} 