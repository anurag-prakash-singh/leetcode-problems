package main

import (
	"fmt"
)

func reverse(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func nextPermutation(nums []int) {
	replacementOffset := -1

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			replacementOffset = i

			break
		}
	}

	if replacementOffset == -1 {
		reverse(nums[replacementOffset+1:])

		return
	}

	nextLargerOffset := replacementOffset + 1
	k := nextLargerOffset

	for ; k < len(nums); k++ {
		if nums[k] <= nums[replacementOffset] {
			nextLargerOffset = k - 1

			break
		}
	}

	if k == len(nums) {
		nextLargerOffset = k - 1
	}

	nums[replacementOffset], nums[nextLargerOffset] = nums[nextLargerOffset], nums[replacementOffset]

	reverse(nums[replacementOffset+1:])
}

func main() {
	inputs := [][]int{{1, 2, 3}, {3, 2, 1}, {1, 1, 5}, {5, 7, 6, 4, 3}, {5, 7, 6, 5, 3}, {1, 3, 2}, {2, 3, 1}}

	for _, input := range inputs {
		fmt.Printf("input: %v; ", input)
		nextPermutation(input)
		fmt.Printf("nextPermutation: %v\n", input)
	}
}
