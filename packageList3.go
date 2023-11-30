package main

import (
	"container/list"
	"fmt"
)

func main() {

	var animals list.List

	animals.PushFront("antelope")
	animals.PushFront("baboon")
	animals.PushFront("Camel")

	for anim := animals.Front(); anim != nil; anim = anim.Next() {
		fmt.Println(anim.Value.(string))
	}

}
