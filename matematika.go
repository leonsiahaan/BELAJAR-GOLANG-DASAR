package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5
	var c = a + b*d
	fmt.Println(c)

	var i = 40
	i += 5 // = i = i + 5
	fmt.Println(i)

	i -= 25 // = i = i - 25
	fmt.Println(i)

	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)

	j-- // j = j - 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)
}
