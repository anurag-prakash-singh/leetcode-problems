package main

import "fmt"

func fillMN(matrix [][]int, r, c, rows, cols, serial int) int {
	if rows <= 0 || cols <= 0 {
		return serial
	}

	if rows == 1 {
		// This is a square matrix (of size 1 x 1) so we can return right away
		matrix[r][c] = serial
		serial++

		return serial
	}

	// First row
	for j := c; j < c+cols; j++ {
		matrix[r][j] = serial
		serial++
	}

	// last column
	for i := r + 1; i < r+rows; i++ {
		matrix[i][c+cols-1] = serial
		serial++
	}

	// last row ... backwards
	for j := c + cols - 2; j > c; j-- {
		matrix[r+rows-1][j] = serial
		serial++
	}

	// first column ... bottom to top
	for i := r + rows - 1; i > r; i-- {
		matrix[i][c] = serial
		serial++
	}

	return serial
}

// func spiralOrder(matrix [][]int) []int {
// 	if len(matrix) == 0 {
// 		return []int{}
// 	}

// 	rows, cols := len(matrix), len(matrix[0])
// 	rowSize, colsSize := len(matrix), len(matrix[0])
// 	result := make([]int, 0, rows*cols)

// 	for i, j := 0, 0; i < rowSize && j < colsSize && rows > 0 && cols > 0; i, j = i+1, j+1 {
// 		result = append(result, listMN(matrix, i, j, rows, cols)...)

// 		rows, cols = rows-2, cols-2
// 	}

// 	return result
// }

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	rows, cols := n, n
	serial := 1

	for s := 0; s <= (n-1)/2; s++ {
		serial = fillMN(matrix, s, s, rows, cols, serial)
		rows, cols = rows-2, cols-2
	}

	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, elem := range row {
			fmt.Printf("%d\t", elem)
		}

		fmt.Println()
	}
}

func test1() {
	matrix := generateMatrix(3)

	printMatrix(matrix)
}

func test2() {
	printMatrix(generateMatrix(1))
}

func test3() {
	printMatrix(generateMatrix(4))
}

func main() {
	// test1()
	// test3()
	test2()
}
