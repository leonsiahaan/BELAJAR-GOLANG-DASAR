package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello	
}

func main() {
	result := getHello("Leon")
	fmt.Println(result)

	fmt.Println(getHello("Siahaan"))
	fmt.Println(getHello("Sitinjak"))
}