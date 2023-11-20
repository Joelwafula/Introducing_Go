package main

import "fmt"

// this is a way of adding elements to an array and returning a new array
func main() {
	slice1 := []int{1, 2, 3}

	slice2 := append(slice1, 4, 5, 6)

	fmt.Println(slice2, slice1)
}
