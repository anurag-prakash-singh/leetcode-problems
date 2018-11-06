package main

import (
	"fmt"
	"sort"
)

func checkIfTriplePresent(result [][]int, triple []int) bool {
	for _, t := range result {
		if t[0] > triple[0] {
			return false
		}

		if t[0] == triple[0] && t[1] == triple[1] && t[2] == triple[2] {
			return true
		}
	}

	return false
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0 /*len(nums)*/, 10)

	for k := 0; k <= len(nums)-3; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		if nums[k] > 0 {
			break
		}

		i := k + 1
		j := len(nums) - 1

		for i < j {
			ijSum := nums[i] + nums[j]

			if ijSum == -nums[k] {
				// add to result
				resultTriple := []int{nums[k], nums[i], nums[j]}

				result = append(result, resultTriple)

				i++
				j--

				for i < j && nums[i] == nums[i-1] {
					i++
				}

				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else if ijSum < -nums[k] {
				i++
			} else {
				j--
			}
		}
	}

	return result
}

func main() {

	test1 := []int{-1, 0, 1, 2, -1, -4}
	ans1 := threeSum(test1)

	fmt.Printf("ans1: %v\n", ans1)

}
