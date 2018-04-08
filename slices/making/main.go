package main

import "fmt"

func main() {
	greeting := []string{
		"Good morning!",
		"Bounjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamate pagi!",
		"Gutten morgen!",
	}

	fmt.Print("[1:4] ")
	fmt.Println(greeting[1:4])
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2])
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:])
	fmt.Print("[:] ")
	fmt.Println(greeting[:])

	for i, entry := range greeting {
		fmt.Println(i, entry)
	}
}
