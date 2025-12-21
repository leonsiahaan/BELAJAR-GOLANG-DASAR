package main

import (
	"BELAJAR-GOLANG-DASAR/database"
	_ "BELAJAR-GOLANG-DASAR/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
