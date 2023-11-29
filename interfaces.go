package main

import "fmt"

// declaring an interface to store the methods
type Shape interface {
	area() float64
}

// declaring a struct of type rectangle
type Rectangle struct {
	length, width float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

//then, we first access the method of the interface

func calculate(s Shape) {
	fmt.Println("area is :", s.area())
}

func main() {
	a := Rectangle{
		length: 10.2,
		width:  22,
	}
	calculate(a)

}
