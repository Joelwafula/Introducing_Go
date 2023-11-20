package main

//defer: this is called after the function completes
import "fmt"

func f1() {
	fmt.Println("first")
}
func f2() {
	fmt.Println("second")
}

func main() {
	defer f1()
	f2()

}

//here, f2() first completes, then followed by the other function
