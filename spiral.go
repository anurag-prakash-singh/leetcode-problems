package main

import "fmt"

func listMN(matrix [][]int, r, c, rows, cols int) []int {
	result := make([]int, 0, 2)

	if rows <= 0 || cols <= 0 {
		return result
	}

	// First row
	for j := c; j < c+cols; j++ {
		result = append(result, matrix[r][j])
	}

	if rows == 1 {
		return result
	}

	// last column
	for i := r + 1; i < r+rows; i++ {
		result = append(result, matrix[i][c+cols-1])
	}

	// last row ... backwards
	for j := c + cols - 2; j > c; j-- {
		result = append(result, matrix[r+rows-1][j])
	}

	if cols == 1 {
		return result
	}

	// first column ... bottom to top
	for i := r + rows - 1; i > r; i-- {
		result = append(result, matrix[i][c])
	}

	return result
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	rows, cols := len(matrix), len(matrix[0])
	rowSize, colsSize := len(matrix), len(matrix[0])
	result := make([]int, 0, rows*cols)

	for i, j := 0, 0; i < rowSize && j < colsSize && rows > 0 && cols > 0; i, j = i+1, j+1 {
		result = append(result, listMN(matrix, i, j, rows, cols)...)

		rows, cols = rows-2, cols-2
	}

	return result
}

func test1() {
	mat := [][]int{
		[]int{5, 2, 13, 15},
		[]int{1, 4, 3, 14},
		[]int{9, 8, 6, 12},
		[]int{11, 10, 7, 16},
	}

	fmt.Printf("spiral order: %v\n", spiralOrder(mat))
}

func test2() {
	mat := [][]int{
		[]int{5, 2, 13, 15},
		[]int{9, 8, 6, 12},
		[]int{11, 10, 7, 16},
	}

	fmt.Printf("spiral order: %v\n", spiralOrder(mat))
}

func test3() {
	mat := [][]int{
		[]int{5, 2, 13, 15},
		[]int{11, 10, 7, 16},
	}

	fmt.Printf("spiral order: %v\n", spiralOrder(mat))
}

func test4() {
	mat := [][]int{
		[]int{5},
	}

	fmt.Printf("spiral order: %v\n", spiralOrder(mat))
}

func test5() {
	mat := [][]int{}

	fmt.Printf("spiral order: %v\n", spiralOrder(mat))
}

func test6() {
	mat := [][]int{
		[]int{3},
		[]int{2},
	}

	fmt.Printf("spiral order: %v\n", spiralOrder(mat))
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
