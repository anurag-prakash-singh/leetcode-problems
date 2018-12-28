package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	gridSum := make([][]int, m)

	for i := 0; i < m; i++ {
		gridSum[i] = make([]int, n)

		for j := 0; j < n; j++ {
			gridSum[i][j] = -1
		}
	}

	gridSum[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				if gridSum[i][j] == -1 {
					gridSum[i][j] = grid[i][j] + gridSum[i-1][j]
				} else {
					gridSum[i][j] = min(gridSum[i][j], grid[i][j]+gridSum[i-1][j])
				}
			}

			if j > 0 {
				if gridSum[i][j] == -1 {
					gridSum[i][j] = grid[i][j] + gridSum[i][j-1]
				} else {
					gridSum[i][j] = min(gridSum[i][j], grid[i][j]+gridSum[i][j-1])
				}
			}
		}
	}

	return gridSum[m-1][n-1]
}

func test1() {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}

	fmt.Printf("minSum: %d\n", minPathSum(grid))
}

func test2() {
	grid := [][]int{
		[]int{1, 3},
		[]int{5, 2},
	}

	fmt.Printf("minSum: %d\n", minPathSum(grid))
}

func main() {
	test1()
	test2()
}
