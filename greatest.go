package main

import "fmt"

func greatestNum(nums ...int) int {

	large := nums[0]

	for i := 0; i < len(nums); i++ {
		if large < nums[i] {
			large = nums[i]
		}
	}
	return large

	// for _, a := range nums {

	// 	if a > nums[] {
	// 		a = nums[]
	// 	}
	// 	return a
	// }
	// return a

}
func main() {
	b := greatestNum(32, 32, 43, 34, 34, 54, 1000)
	fmt.Println(b)
}
