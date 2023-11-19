package main

import "fmt"

func main() {
	var i int = 32
	switch {
	case i == 0:
		fmt.Println("zero")
	case i == 1:
		fmt.Println("one")
	case i == 2:
		fmt.Println("two")

	default:
		fmt.Println("there is no such case")
	}
}
