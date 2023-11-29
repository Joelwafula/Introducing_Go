package main

import "fmt"

func factorial2(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial2(x-1)
}
func main() {
	b := factorial2(5)
	fmt.Println(b)
}
