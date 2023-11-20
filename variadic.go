package main

import "fmt"

// we use this to initalise multiple functions
func add(args ...int) int {

	total := 0

	for _, v := range args {
		total += v

	}
	return total

}

func main() {
	a := add(12, 12, 32, 13, 54)

	fmt.Println(a)
}
