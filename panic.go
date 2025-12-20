package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func runApp(error bool) {
	defer endApp()
	
	if error {
		panic("App Eror")
	}

}

func main() {
	runApp(true)
	// runApp(false)
	fmt.Println("Leon Siahaan")
}