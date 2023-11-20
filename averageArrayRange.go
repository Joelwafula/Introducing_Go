package main

import "fmt"

func main() {
	x := [4]int{12, 13, 14, 15}

	var sum int

	for _, num := range x {

		sum += num

	}
	fmt.Println(sum / len(x))
}
