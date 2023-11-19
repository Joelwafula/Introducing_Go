package main

import "fmt"

func main() {

	fmt.Println("Enter a positive integer :")
	var input uint64

	fmt.Scanf("%d", &input)

	output := input * 2

	fmt.Println(output)
}
