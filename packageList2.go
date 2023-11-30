package main

import (
	"container/list"
	"fmt"
)

func main() {
	var myList list.List

	myList.PushBack("Joel")
	myList.PushBack("Wafula")
	myList.PushBack("Simiyu")

	for element := myList.Front(); element != nil; element = element.Next() {

		fmt.Println(element.Value.(string))

	}
}
