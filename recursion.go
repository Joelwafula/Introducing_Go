//this is where a function is able to call itself

package main

import "fmt"

func factorial(x uint) uint {

	if x == 0 {
		return 1
	}
	return x * factorial(x-1)

}

func main() {
	a := factorial(4)
	fmt.Println(a)

}
