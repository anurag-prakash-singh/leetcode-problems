package main

import (
	"fmt"
)

/*
Cases:

1. a < b < c -- do nothing
2. a < b > c
	2.1 a <= c -- set b = c
	2.2 a > c  -- can't fix 
3. a < b = c -- do nothing
4. a > b < c
	4.1 a <= c -- set b = c
	4.2 a  > c -- can't fix
5. a > b > c -- can't fix
6. a > b = c -- can't fix
7. a = b < c -- do nothing
8. a = b > c -- can't fix
9. a = b = c -- do noting

*/

func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	a, b, c := 0, 0, 0
	numChanges := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			b, c = nums[i], nums[i + 1]
			
			if b > c {
				numChanges++
				b = c
				nums[i] = b
			}

			continue
		}

		if i == len(nums) - 1 {
			a, b = nums[i - 1], nums[i]

			if a > b {
				numChanges++
				b = a
				nums[i] = b
			}

			continue
		}

		a, b, c = nums[i-1], nums[i], nums[i + 1]

		// 2. a < b > c
		// 	2.1 a <= c -- set b = c
		// 	2.2 a > c  -- can't fix 
		if (a < b && b > c) || (a > b && b < c) {
			if a <= c {
				numChanges++
				b = c
				nums[i] = b
			}

			continue
		}
	}

	// Check if non-decreasing
	for i := 0; i < len(nums); i++ {
		if i < len(nums) - 1 {
			if nums[i] > nums[i + 1] {
				return false
			}
		}
	}

	return numChanges <= 1
}

type testcase struct {
	nums []int
	expectedResult bool
}

func tests() {
	tcs := []testcase {
		testcase{nums: []int{4}, expectedResult: true},
		testcase{nums: []int{4,2}, expectedResult: true},
		testcase{nums: []int{4,2,3}, expectedResult: true},
		testcase{nums: []int{4,3,2}, expectedResult: false},
		testcase{nums: []int{1,1,3,2}, expectedResult: true},
		testcase{nums: []int{1,2,3,2}, expectedResult: true},
		testcase{nums: []int{1,2,3,2,5}, expectedResult: true},
		testcase{nums: []int{1,7,3,2,5}, expectedResult: false},
		testcase{nums: []int{1,7,3,2,5,5}, expectedResult: false},
		testcase{nums: []int{1,2,4,5,3}, expectedResult: true},
	}

	for i, tc := range tcs {
		result := checkPossibility(tc.nums)

		if result == tc.expectedResult {
			fmt.Printf("TC %d PASSED\n", i)
		} else {
			fmt.Printf("TC %d FAILED; expected: %v, actual: %v\n", i, tc.expectedResult, result)
		}
	}
}

func main() {
	tests()
}