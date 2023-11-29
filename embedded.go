package main

import "fmt"

type Person struct {
	name string
}

func (p Person) Talk() {

	fmt.Println("Hello your name is :", p.name)

}
func main() {
	a := Person{
		name: "Joel",
	}
	a.Talk()
}
