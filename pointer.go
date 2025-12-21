package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Medan", "Sumatera Utara", "Indonesia"}
	// address2 := address1 // copy value

	// address2.City = "Siantar"

	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2)  // berubah jadi siantar

	var address1 Address = Address{"Medan", "Sumatera Utara", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Siantar"

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah jadi Siantar

}
