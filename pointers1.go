package main

import "fmt"

func zer(xP *int) {
	*xP = 2

}

func main() {
	x := 5
	zer(&x)

	fmt.Println(x)

}
