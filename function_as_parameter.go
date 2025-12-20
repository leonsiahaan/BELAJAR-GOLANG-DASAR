package main

import "fmt"

type Filter func(string) string

func  sayHelloWithName(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello " + filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithName("Leon", spamFilter)

	filter := spamFilter
	sayHelloWithName("anjing", filter)
} 