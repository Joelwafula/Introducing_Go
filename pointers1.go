package main

import "fmt"

func zer(xPtr *int) {
	*xPtr = 0

}

func main() {
	x := 5
	zer(&x)

	fmt.Println(x)

}
