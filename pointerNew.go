package main

import "fmt"

func one(xPt *int) {
	*xPt = 1
}

//we initialise the pointer using new

func main() {
	xPt := new(int)
	one(xPt)
	fmt.Println(&xPt)
	fmt.Println(*xPt)

}
