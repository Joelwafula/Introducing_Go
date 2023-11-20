package main

import "fmt"

func main() {
	x := [6]int{21, 23, 43, 56, 75, 65}

	var sum int

	for i := 0; i < 6; i++ {
		sum += x[i]
	}
	fmt.Println(sum / len(x))

}
