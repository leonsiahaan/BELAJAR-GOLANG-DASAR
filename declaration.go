package main

func main() {
	type NoKTP string

	var ktpLeon NoKTP = "1111111111111"

	var contoh string = "33333333333333"
	var contohKtp NoKTP = NoKTP(contoh)

	println(ktpLeon)
	println(contohKtp)
}
