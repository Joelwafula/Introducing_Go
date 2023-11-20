package main

import "fmt"

func average(xs []float64) float64 {

	total := 0.0

	for _, v := range xs {

		total += v

	}
	return total / float64(len(xs))

}

func main() {

	xs := []float64{10, 20, 23, 232, 32, 23}

	ys := []float64{23, 34, 434, 43, 353, 53.35}
	a := average(xs)
	b := average(ys)

	fmt.Println(a)
	fmt.Println(b)
}
