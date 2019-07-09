package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxKnapsackValue(w int, ws, vs []int) int {
	dp := make([][]int, w+1)
	n := len(ws)

	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	// dp[i][j] is the maximum value of the knapsack containing
	// a subset of the items [0..(j - 1)] whose weights add up to
	// less than or equal to i
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= w; i++ {
		dp[i][0] = 0

		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i][j-1]

			if i-ws[j-1] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-ws[j-1]][j-1]+vs[j-1])
			}
		}
	}

	return dp[w][n]
}

type testcase struct {
	weights          []int
	values           []int
	w                int
	expectedMaxValue int
}

func test() {
	tcs := []testcase{
		testcase{weights: []int{3, 2, 4, 1}, values: []int{100, 20, 60, 40}, w: 5, expectedMaxValue: 140},
		testcase{weights: []int{10, 20, 30}, values: []int{60, 100, 120}, w: 50, expectedMaxValue: 220},
	}

	for i, tc := range tcs {
		result := maxKnapsackValue(tc.w, tc.weights, tc.values)

		if result != tc.expectedMaxValue {
			fmt.Printf("Test case %d FAILED (expected: %d, actual: %d)\n", i, tc.expectedMaxValue, result)
		} else {
			fmt.Printf("Test case %d PASSED\n", i)
		}
	}
}

func main() {
	test()
}
