package main

import "fmt"

func setZeroes(matrix [][]int) {
	zeroFirstCol := false
	m := len(matrix)

	if m == 0 {
		return
	}

	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if j == 0 {
					zeroFirstCol = true

					continue
				} else {
					matrix[i][0] = 0
					matrix[0][j] = 0
				}
			}
		}
	}

	// fmt.Printf("intermed: %v\n", matrix)

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if zeroFirstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

}

func test1() {
	matrix := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}

	setZeroes(matrix)

	fmt.Printf("ans: %v\n", matrix)
}

func test2() {
	matrix := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}

	setZeroes(matrix)

	fmt.Printf("ans: %v\n", matrix)
}

func main() {
	test1()
	test2()
}
