package main

import "fmt"

func main() {
	var x int

	fmt.Println("Enter an integer")

	fmt.Scanf("%d", &x)

	if (x / 2) > 0 {
		fmt.Println(true)

	} else {
		fmt.Println(false)
	}

}
