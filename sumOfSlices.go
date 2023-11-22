package main

import "fmt"

func main(){
	slice := []int{21,23,33,43,43}

	var total int;

	for _, t := range slice {
		total += t
	}
	fmt.Println(total)
}