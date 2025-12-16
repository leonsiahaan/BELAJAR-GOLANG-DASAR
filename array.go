package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Leon"
	names[1] = "Siahaan"
	names[2] = "hinalang"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		80,
		75,
		85,
		95,
	}
	fmt.Println(values)
	fmt.Println(len(values))

	var names2 = [3]string{
		"Leon",
		"Siahaan",
	}
	fmt.Println(names2)
	fmt.Println(names2[0])
	fmt.Println(names2[1])
}