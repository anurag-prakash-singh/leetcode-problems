package main

import "fmt"

func binSearch(matrix [][]int, target, low, high, rows, cols int) bool {
	if low > high {
		return false
	}

	mid := (low + high) / 2
	midRow := mid / cols
	midCol := mid % cols

	if matrix[midRow][midCol] == target {
		return true
	}

	if matrix[midRow][midCol] < target {
		return binSearch(matrix, target, mid+1, high, rows, cols)
	}

	return binSearch(matrix, target, low, mid-1, rows, cols)
}

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)

	if rows == 0 {
		return false
	}

	cols := len(matrix[0])

	return binSearch(matrix, target, 0, rows*cols-1, rows, cols)
}

func test1() {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}

	fmt.Printf("ans: %v\n", searchMatrix(matrix, 3))
}

func test2() {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}

	fmt.Printf("ans: %v\n", searchMatrix(matrix, 13))
}

func test3() {
	matrix := [][]int{
		[]int{1, 3, 5, 7},
	}

	fmt.Printf("ans: %v\n", searchMatrix(matrix, 5))
	fmt.Printf("ans: %v\n", searchMatrix(matrix, 7))
	fmt.Printf("ans: %v\n", searchMatrix(matrix, 42))
}

func test4() {
	matrix := [][]int{
		[]int{1},
	}

	fmt.Printf("ans: %v\n", searchMatrix(matrix, 3))
	fmt.Printf("ans: %v\n", searchMatrix(matrix, 1))
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
