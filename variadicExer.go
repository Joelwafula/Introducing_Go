package main

import "fmt"

func largerstNumber(num ...int) int {

	largest := 0

	for _, large := range num {

		if large > largest {
			largest = large
		}
		fmt.Println(largest)

	}
	return largest
}
func main() {
	a := largerstNumber(21, 42, 35, 53, 535, 24, 24)
	fmt.Println(a)

}
