package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxLoot := nums[0]

	for i, _ := range nums {
		if i == 0 {
			continue
		}

		j := i - 2
		k := i - 3

		if j >= 0 {
			if k >= 0 {
				nums[i] = max(nums[i]+nums[j], nums[i]+nums[k])
			} else {
				nums[i] += nums[j]
			}

			maxLoot = max(maxLoot, nums[i])
		} else {
			maxLoot = max(nums[i], nums[i-1])
		}
	}

	return maxLoot
}

func test1() {
	tests := [][]int{
		[]int{1, 2, 3, 1}, []int{4},
		[]int{2, 7, 9, 3, 1}, []int{12},
		[]int{5, 5, 10, 100, 10, 5}, []int{110},
		[]int{1, 2, 3}, []int{4},
		[]int{10, 1}, []int{10},
		[]int{10}, []int{10},
		[]int{1, 10}, []int{10},
		[]int{3, 2, 1}, []int{4},
	}

	for i := 0; i < len(tests); i += 2 {
		result := rob(tests[i])

		if result == tests[i+1][0] {
			fmt.Printf("Test case %d PASSED\n", i/2)
		} else {
			fmt.Printf("Test case %d FAILED. Expected: %d, got %d\n", i/2, tests[i+1][0], result)
		}
	}
}

func main() {
	test1()
}
