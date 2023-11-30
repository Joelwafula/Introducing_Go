package main

//sort is used for soring arbitrary data
import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}
type ByName []Person

func (ps ByName) Len() int {
	return len(ps)

}
func (ps ByName) Less(i, j int) bool {
	return ps[i].name < ps[j].name

}

func (ps ByName) Swap()
