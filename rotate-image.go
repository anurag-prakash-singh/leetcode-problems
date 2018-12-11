package main

import (
	"fmt"
)

func reverse(arr []int) {
	n := len(arr)

	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

func rotate(matrix [][]int) {
	n := len(matrix)

	// Flip across top-left/bottom-right diagonal
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse each row
	for i := 0; i < n; i++ {
		reverse(matrix[i])
	}
}

func main() {
	testMat1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	testMat2 := [][]int{
		{1},
	}

	testMat3 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(testMat1)
	rotate(testMat2)
	rotate(testMat3)

	fmt.Printf("mat: %v\n", testMat3)
}
