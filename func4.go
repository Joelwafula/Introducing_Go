//named return type

package main

import "fmt"

func f() (int, int) {
	return 34, 35

}

func main() {

	_, y := f()

	fmt.Println(y)

}
