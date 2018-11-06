package main

import (
	"fmt"
	"math"
	"sort"
)

func intAbs(n int) int {
	return (int)(math.Abs((float64)(n)))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]

	for k := 0; k <= len(nums)-3; k++ {
		i := k + 1
		j := len(nums) - 1

		if nums[k] > 0 && nums[k] > target {
			return result
		}

		for i < j {
			ijkSum := nums[i] + nums[j] + nums[k]

			if ijkSum == target {
				return target
			} else {
				if intAbs(target-ijkSum) < intAbs(target-result) {
					result = ijkSum
				}

				if ijkSum < target {
					i++
				} else {
					j--
				}
			}
		}
	}

	return result
}

func main() {

	test1 := []int{-1, 2, 1, -4}
	ans1 := threeSumClosest(test1, 1)

	fmt.Printf("ans1: %v\n", ans1)

}
