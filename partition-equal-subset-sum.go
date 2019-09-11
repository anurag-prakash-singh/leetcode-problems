package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	if sum%2 != 0 {
		return false
	}

	halfSum := sum / 2
	part := make([][]bool, halfSum+1)
	n := len(nums)

	for i := 0; i <= halfSum; i++ {
		part[i] = make([]bool, n+1)
	}

	for i := 0; i <= n; i++ {
		part[0][i] = true
	}

	for s := 1; s <= halfSum; s++ {
		part[s][0] = false

		for j := 1; j <= n; j++ {
			if part[s][j-1] {
				part[s][j] = true

				continue
			}

			if s-nums[j-1] >= 0 {
				// parts[s][j] is true if there's a subset of nums[0 .. j - 1] that sums to s.
				// In other words, if there's a subset of nums[0 .. j - 2] that sums to s - nums[j - 1].
				part[s][j] = part[s-nums[j-1]][j-1]
			}
		}
	}

	return part[halfSum][n]
}

type testCase struct {
	nums     []int
	expected bool
}

func tests() {
	testCases := []testCase{
		testCase{nums: []int{1, 3, 2}, expected: true},
		testCase{nums: []int{1, 5, 11, 5}, expected: true},
		testCase{nums: []int{1, 2, 3, 5}, expected: false},
	}

	for i, tc := range testCases {
		result := canPartition(tc.nums)

		if result == tc.expected {
			fmt.Printf("Test number %d: PASS\n", i)
		} else {
			fmt.Printf("Test number %d: FAILED (expected: %v, got %v)\n", i, tc.expected, result)
		}
	}
}

func main() {
	tests()
}
