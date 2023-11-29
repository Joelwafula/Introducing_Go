package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, z float64
}

func (c Circle) area() float64 {
	return math.Pi * c.z * c.z
}
func main() {
	c := Circle{0, 0, 10}

	fmt.Println(c.area())

}
