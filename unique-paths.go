package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	grid := make([][]int, m)

	for i, _ := range grid {
		grid[i] = make([]int, n)
	}

	grid[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rightR, rightC := i, j+1
			downR, downC := i+1, j

			if rightC < n {
				grid[rightR][rightC] += grid[i][j]
			}

			if downR < m {
				grid[downR][downC] += grid[i][j]
			}
		}
	}

	return grid[m-1][n-1]
}

type testcase struct {
	m              int
	n              int
	expectedResult int
}

func tests() {
	testcases := []testcase{
		testcase{m: 3, n: 2, expectedResult: 3},
		testcase{m: 7, n: 3, expectedResult: 28},
	}

	for i, tc := range testcases {
		result := uniquePaths(tc.m, tc.n)

		if result != tc.expectedResult {
			fmt.Printf("Test %d FAILED (expected: %d; actual %d)\n", i, tc.expectedResult, result)
		} else {
			fmt.Printf("Test %d PASSED\n", i)
		}
	}
}

func main() {
	tests()
}
