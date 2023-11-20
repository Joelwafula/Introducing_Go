package main

import "fmt"

// function call stack
//import "fmt"

func main() {
	a := f1()
	fmt.Println(a)

}
func f1() int {
	return f2()

}

func f2() int {
	return 1
}
