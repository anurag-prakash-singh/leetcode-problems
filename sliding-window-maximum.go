package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func fillMaxLeftRight(nums, maxL, maxR []int) {
	if len(nums) == 0 {
		return
	}

	// fill maxL
	maxLSoFar := nums[0]

	for i, num := range nums {
		if num >= maxLSoFar {
			maxLSoFar = num
		}

		maxL[i] = maxLSoFar
	}

	// fill maxR
	maxRSoFar := nums[len(nums)-1]

	for i := len(nums) - 1; i >= 0; i -= 1 {
		if nums[i] >= maxRSoFar {
			maxRSoFar = nums[i]
		}

		maxR[i] = maxRSoFar
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	maxL := make([]int, len(nums))
	maxR := make([]int, len(nums))
	n := len(nums)

	for i := 0; i < n; i += k {
		chunkSize := min(n-i, k)

		fillMaxLeftRight(nums[i:i+chunkSize], maxL[i:i+chunkSize], maxR[i:i+chunkSize])
	}

	result := make([]int, n-k+1)

	for i := 0; i < n-k+1; i++ {
		result[i] = max(maxR[i], maxL[i+k-1])
	}

	return result
}

func checkSliceEquality(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

type testCase struct {
	input          []int
	k              int
	expectedResult []int
}

func tests() {
	testCases := []testCase{
		testCase{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		testCase{[]int{4, 3, 5, 6, 1}, 3, []int{5, 6, 6}},
	}

	for i, tc := range testCases {
		result := maxSlidingWindow(tc.input, tc.k)

		if checkSliceEquality(result, tc.expectedResult) {
			fmt.Printf("test case %d PASSED\n", i)
		} else {
			fmt.Printf("test case %d failed. Got %v, expected %v\n", i, result, tc.expectedResult)
		}
	}
}

func main() {
	tests()
}
