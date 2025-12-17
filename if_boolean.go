package main

import "fmt"

func main() {
	name := "Kurniawan"

	if name == "Leon" {
		fmt.Println("Hello Leon")
	} else if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hai Booleh kenalan")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah pas")
	}
}