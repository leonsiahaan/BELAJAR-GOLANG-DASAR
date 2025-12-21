package main

import (
	"BELAJAR-GOLANG-DASAR/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Leon")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.Version)
	fmt.Println(helper.SayGoodbye("Leon"))
}
