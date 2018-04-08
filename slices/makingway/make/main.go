package main

import "fmt"

func main() {
	sliceMake := make([]string, 35)
	sliceMakeMd := make([][]string, 35)

	fmt.Println(sliceMake)
	fmt.Println(sliceMakeMd)
	fmt.Println(sliceMake == nil)
}
