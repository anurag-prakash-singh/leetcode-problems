package main

import "fmt"

func sortedSquares(A []int) []int {
	i, j, k := 0, len(A)-1, len(A)-1
	result := make([]int, len(A))

	for i <= j {
		if A[i]*A[i] > A[j]*A[j] {
			A[k] = A[i] * A[i]
			i++
			k--

			continue
		}

		result[k] = A[j] * A[j]
		j--
		k--
	}

	return result
}

func runTests() {
	fmt.Printf("sortedSquares: %v\n", sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Printf("sortedSquares: %v\n", sortedSquares([]int{}))
	fmt.Printf("sortedSquares: %v\n", sortedSquares([]int{1}))
	fmt.Printf("sortedSquares: %v\n", sortedSquares([]int{-1}))
	fmt.Printf("sortedSquares: %v\n", sortedSquares([]int{-7, -3, 2, 3, 11}))
}

func main() {
	runTests()
}
