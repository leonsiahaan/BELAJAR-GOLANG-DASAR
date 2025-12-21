package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	var address1 Address = Address{"Medan", "Sumatera Utara", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Siantar"

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah jadi Siantar

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // tetap siantar
	fmt.Println(address2) // berubah jakarta
}
