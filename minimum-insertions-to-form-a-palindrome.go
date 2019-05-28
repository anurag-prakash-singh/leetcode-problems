package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func findMinInsertions(str string) int {
	chars := []rune(str)
	n := len(str)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
	}

	for l := 2; l <= n; l++ {
		for lo := 0; lo <= n-l; lo++ {
			hi := lo + l - 1

			if chars[lo] == chars[hi] {
				if hi == lo+1 {
					dp[lo][hi] = 0
				} else {
					dp[lo][hi] = dp[lo+1][hi-1]
				}
			} else {
				dp[lo][hi] = 1 + min(dp[lo+1][hi], dp[lo][hi-1])
			}
		}
	}

	return dp[0][n-1]
}

type testCase struct {
	str            string
	expectedResult int
}

func tests() {
	testCases := []testCase{
		testCase{str: "ab", expectedResult: 1},
		testCase{str: "aa", expectedResult: 0},
		testCase{str: "abcd", expectedResult: 3},
		testCase{str: "abcda", expectedResult: 2},
		testCase{str: "abcde", expectedResult: 4},
	}

	for i, tc := range testCases {
		result := findMinInsertions(tc.str)

		if result != tc.expectedResult {
			fmt.Printf("Test %d FAILED. Expected %d, got %d\n", i, tc.expectedResult, result)
		} else {
			fmt.Printf("Test %d PASSED\n", i)
		}
	}
}

func main() {
	tests()
}
