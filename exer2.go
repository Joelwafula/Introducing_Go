/*Write a program that prints the numbers from 1 to 100, but for multiples of
three, print “Fizz” instead of the number, and for the multiples of five, print
“Buzz.” For numbers that are multiples of both three and five, print “FizzBuzz.”*/

package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println("buzz")
		}
		if i%5 == 0 {
			fmt.Println("fizz")
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")

		}
	}
}
