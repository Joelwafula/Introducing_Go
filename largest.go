package main

import "fmt"

func main() {
	x := []int{
		100, 96, 600, 68,
		57, 82, 63, 70,
		43, 34, 200, 27,
		19, 1900, 9, 17,
	}

	largest := x[0]

	for _, num := range x {
		if num > largest {
			largest = num
		}
	}
	fmt.Println(largest)
}
