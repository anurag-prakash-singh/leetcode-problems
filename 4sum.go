package main

import (
	"fmt"
	"sort"
)

// nums is assumed to be pre-sorted
func threeSum(nums []int, target int) [][]int {
	result := make([][]int, 0, 10)

	for k := 0; k <= len(nums)-3; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i := k + 1
		j := len(nums) - 1

		for i < j {
			ijkSum := nums[i] + nums[j] + nums[k]

			if ijkSum == target {
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
			} else if ijkSum < target {
				i++
			} else {
				j--
			}
		}
	}

	return result
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0, 10)

	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if i > len(nums)-4 {
			break
		}

		fmt.Printf("v: %d; looking for %d in %v\n", v, target-v, nums[i+1:])
		resultsWOv := threeSum(nums[i+1:], target-v)

		for _, r := range resultsWOv {
			result = append(result, []int{v, r[0], r[1], r[2]})
		}
	}

	return result
}

func main() {

	test1 := []int{1, -2, -5, -4, -3, 3, 3, 5}
	ans1 := fourSum(test1, -11)
	fmt.Printf("ans1: %v\n", ans1)

}
