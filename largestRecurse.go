package main

import "fmt"

func getLargest(A []int, n int) int {
	if n == 0 {
		return A[0]
	}
	return max(A[n-1], getLargest(A[:n-1], n-1))

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func main() {
	A := []int{1, 4, 45, 6, -50, 10, 2}
	n := len(A)

	// Function calling
	fmt.Println(getLargest(A, n))

}
