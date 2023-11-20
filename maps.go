//map is an unordered selection of key value pairs

package main

import "fmt"

func main() {

	x := make(map[int]string)
	x[1] = "Joel"
	x[2] = "Wafula"
	x[3] = "simiyu"

	fmt.Println(x[1])

}
