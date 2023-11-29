package main

import (
	//"encoding/base32"
	"fmt"
)

type Shp interface {
	areas() float32
}

type Cuboid struct {
	len, wid, het float32
}
type Triangle struct {
	base, height float32
}

func (c Cuboid) areas() float32 {
	return c.het * c.len * c.wid
}
func (t Triangle) areas() float32 {
	return 0.5 * t.base * t.height
}

func getArea(g Shp) float32 {
	return g.areas()
}

func main() {
	r := Triangle{
		base:   20,
		height: 30,
	}
	q := Cuboid{
		het: 21,
		len: 32,
		wid: 54,
	}
	a := getArea(r)
	b := getArea(q)

	fmt.Println(a, b)

}
