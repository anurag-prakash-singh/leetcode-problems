package main

import (
	"fmt"
)

func max2(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func largestRectangleArea(heights []int) int {
	// heightStack := make([]int, 0, len(heights)+1)
	indexStack := make([]int, 0, len(heights)+1)
	maxArea := 0

	indexStack = append(indexStack, -1)

	for i := 0; i < len(heights); i++ {
		// We have an inversion
		for indexStack[len(indexStack)-1] != -1 && heights[indexStack[len(indexStack)-1]] >= heights[i] {
			h := heights[indexStack[len(indexStack)-1]]
			indexStack = indexStack[:len(indexStack)-1]

			candidateArea := h * (i - indexStack[len(indexStack)-1] - 1)
			maxArea = max2(maxArea, candidateArea)
		}

		indexStack = append(indexStack, i)
	}

	// We may still have some bars left in our stack - this will happen if there was no
	// inversion at the end. Process them.
	for indexStack[len(indexStack)-1] != -1 {
		h := heights[indexStack[len(indexStack)-1]]
		indexStack = indexStack[:len(indexStack)-1]

		candidateArea := h * (len(heights) - indexStack[len(indexStack)-1] - 1)
		maxArea = max2(maxArea, candidateArea)
	}

	return maxArea
}

type testCase struct {
	heights      []int
	expectedArea int
}

func runTests() {
	testCases := []testCase{
		testCase{heights: []int{1, 3, 2, 1, 2}, expectedArea: 5},
		testCase{heights: []int{4}, expectedArea: 4},
		testCase{heights: []int{2, 1, 5, 6, 2, 3}, expectedArea: 10},
		testCase{heights: []int{1, 3, 2}, expectedArea: 4},
	}

	for i, tc := range testCases {
		result := largestRectangleArea(tc.heights)

		if result != tc.expectedArea {
			fmt.Printf("Testcase %d FAILED; expected %d, got %d\n", i, tc.expectedArea, result)
		} else {
			fmt.Printf("Testcase %d PASSED\n", i)
		}
	}
}

func main() {
	runTests()
}
