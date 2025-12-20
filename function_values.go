package main

import "fmt"

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	misal := getGoodbye
	contoh := getGoodbye

	fmt.Println(misal("Leon"))
	fmt.Println(contoh("Siahaan"))
}