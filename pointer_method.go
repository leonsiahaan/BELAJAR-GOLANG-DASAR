package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	leon := Man{"Leon"}
	leon.Married()

	fmt.Println(leon.Name)
}
