package main

//first we will use itteration, where we compare two arrays at a time
import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	small := x[0]

	for i := 1; i < len(x); i++ {
		if x[i] < small {
			small = x[i]

		}

	}
	fmt.Println(small)
}
