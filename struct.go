package main

import "fmt"

type Customer struct {
	Name, Address string
	Age 		  int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var Leon Customer
	fmt.Println(Leon)

	Leon.Name = "Leon Siahaan"
	Leon.Address = "Jakarta"
	Leon.Age = 25
	fmt.Println(Leon)
	fmt.Println(Leon.Name)
	fmt.Println(Leon.Address)
	fmt.Println(Leon.Age)

	andi := Customer{
		Name:    "Andi",
		Address: "Bandung",
		Age:     30,
	}
	fmt.Println(andi)

	budi := Customer{"Budi", "Surabaya", 35}
	fmt.Println(budi)

	Leon.sayHello("Amir")
	andi.sayHello("Sitinjak")
	budi.sayHello("Budi")
	
} 
