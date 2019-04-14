package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) == 0 || k == 0 {
		return 0
	}

	sChars := []rune(s)
	start, end := -1, -1 // start is exclusive: (start, end]
	maxLen := 0
	currDistinctCount := 0
	occMap := make(map[rune]int)

	for end < len(sChars) {
		if currDistinctCount <= k {
			end++

			if end >= len(sChars) {
				break
			}

			occMap[sChars[end]]++

			if occMap[sChars[end]] == 1 {
				currDistinctCount++
			}

			if currDistinctCount <= k {
				maxLen = max(maxLen, end-start)
			}
		} else {
			start++

			occMap[sChars[start]] -= 1

			if occMap[sChars[start]] == 0 {
				currDistinctCount--
			}
		}
	}

	return maxLen
}

type testCase struct {
	s              string
	k              int
	expectedResult int
}

func test() {
	tests := []testCase{
		testCase{"eceba", 2, 3},
		testCase{"aa", 1, 2},
		testCase{"abcba", 2, 3},
		testCase{"abcd", 1, 1},
	}

	for i, tc := range tests {
		result := lengthOfLongestSubstringKDistinct(tc.s, tc.k)

		if result == tc.expectedResult {
			fmt.Printf("Test case %d PASSED\n", i)
		} else {
			fmt.Printf("Test case %d FAILED. Got %d but expected %d\n", i, result, tc.expectedResult)
		}
	}
}

func main() {
	test()
}
