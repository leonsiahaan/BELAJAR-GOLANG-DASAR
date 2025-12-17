package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Leon":	
		fmt.Println("Hello Leon")
	case "Eko":
		fmt.Println("Hello Eko")
	default: 
		fmt.Println("Hai Booleh kenalan")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah pas")
	}

	name = "Kurniawan"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah pas")
	}
}

