package main

import "fmt"

func max2(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	sum := 0
	end := 0

	for end < len(nums) {
		if nums[end]+sum < 0 {
			maxSum = max2(maxSum, sum+nums[end])
			end++
			sum = 0

			continue
		}

		maxSum = max2(maxSum, sum+nums[end])
		sum += nums[end]
		end++
	}

	return maxSum
}

func test1() {
	tc := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Printf("maxSum = %d\n", maxSubArray(tc))
}

func test2() {
	tc := []int{-3, -2, 1}

	fmt.Printf("maxSum = %d\n", maxSubArray(tc))
}

func test3() {
	tc := []int{4, -1, 2, 1}

	fmt.Printf("maxSum = %d\n", maxSubArray(tc))
}

func main() {
	test1()
	test2()
	test3()
}
