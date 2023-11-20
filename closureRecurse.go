package main

//this is a function that returns another function as a paramater and its returned in the code
import "fmt"

func makeEvenGenerator() func() uint {

	i := uint(0)

	return func() (j uint) {

		j = i

		i += 2
		return j

	}

}
func main() {
	nextEven := makeEvenGenerator()

	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

}
