package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	if len(nums) == 1 && nums[0] == 1 {
		return 2
	}

	max := 0

	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	if max <= 0 {
		return 1
	}

	for i, n := range nums {
		if n <= 0 {
			nums[i] = max
		}
	}

	nums = append(nums, max)

	for _, n := range nums {
		n = abs(n)
		if n > 0 && n < len(nums) && nums[n] > 0 {
			nums[n] *= -1
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}

	return len(nums)
}

func test1() {
	testCases := [][]int{
		[]int{0, 1, 2, 3, 4}, []int{5},
		[]int{-1}, []int{1},
		[]int{0}, []int{1},
		[]int{1}, []int{2},
		[]int{2}, []int{1},
		[]int{0, 0}, []int{1},
		[]int{0, -1}, []int{1},
		[]int{1, 2, 0}, []int{3},
		[]int{3, 4, -1, 1}, []int{2},
		[]int{7, 8, 9, 11, 12}, []int{1},
	}

	for i := 0; i < len(testCases); i += 2 {
		nums, expResult := testCases[i], testCases[i+1][0]
		result := firstMissingPositive(nums)

		if expResult == result {
			fmt.Printf("PASS: fmp(%v) == %d\n", nums, expResult)
		} else {
			fmt.Printf("FAILED: expected fmp(%v) == %d, but got %d\n", nums, expResult, result)
		}
	}
}

func main() {
	test1()
}
