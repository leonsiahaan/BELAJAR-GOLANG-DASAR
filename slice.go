package main

import "fmt"

func main() {
	names := [...]string{"Leon", "Siahaan", "Nadia", "Simatupang", "Doraemon", "Nobita", "Suneo", "Giant"}

	slice1 := names[4:6]
	fmt.Println(slice1)
	slice2 := names[:3]
	fmt.Println(slice2)
	slice3 := names[3:]
	fmt.Println(slice3)
	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	
	dayslice1 := days[5:] // sabtu, minggu
	fmt.Println(dayslice1)

	dayslice1[0] = "Sabtu Baru"
	dayslice1[1] = "Minggu Baru"
	fmt.Println(dayslice1)
	fmt.Println(days)

	dayslice2 := append(dayslice1, "Libur Baru")
	dayslice2[0] = "Sabtu lama"
	//daysbaru := [...] string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu baru", "minggu baru", "libur baru"})
	fmt.Println(dayslice1)
	fmt.Println(dayslice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Leon"
	newSlice[1] = "Siahaan"
	
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Nadia")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Nadia"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)	

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlace := []int{1, 2, 3, 4, 5}
	
	fmt.Println(iniArray)
	fmt.Println(iniSlace)
	
}