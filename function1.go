package main

import (
	"fmt"
)

func main() {
	xs := []int{98, 98, 97, 95, 94}

	total := 0

	for _, v := range xs {
		total += v

	}
	fmt.Println(total / len(xs))
}
