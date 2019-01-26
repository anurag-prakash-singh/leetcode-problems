package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return computePowerSetWithDup(nums, 0, []int{}, [][]int{})
}

func computePowerSetWithDup(nums []int, pos int, setSoFar []int, results [][]int) [][]int {
	if pos >= len(nums) {
		resultSet := make([]int, len(setSoFar))
		copy(resultSet, setSoFar)
		return append(results, resultSet)
	}

	posrlc := 0

	for i := pos; i < len(nums) && nums[i] == nums[pos]; i++ {
		posrlc++
	}

	// Case 1: move forward, excluding nums[pos]
	tempSetWithoutPos := make([]int, len(setSoFar))
	copy(tempSetWithoutPos, setSoFar)

	newResults := computePowerSetWithDup(nums, pos+posrlc, tempSetWithoutPos, results)

	// Case 2: move forwawrd, including nums[pos]
	for i := 1; i <= posrlc; i++ {
		tempSetWithPos := []int{}

		tempSetWithPos = append(tempSetWithPos, setSoFar...)

		for j := 0; j < i; j++ {
			tempSetWithPos = append(tempSetWithPos, nums[pos])
		}

		newResults = computePowerSetWithDup(nums, pos+posrlc, tempSetWithPos, newResults)
	}

	return newResults
}

func test1() {
	results := subsetsWithDup([]int{1, 2, 3})

	fmt.Printf("power set: %v\n", results)
}

func test2() {
	results := subsetsWithDup([]int{})

	fmt.Printf("power set: %v\n", results)
}

func test3() {
	results := subsetsWithDup([]int{1})

	fmt.Printf("power set: %v\n", results)
}

func test4() {
	results := subsetsWithDup([]int{1, 2, 2})

	fmt.Printf("power set: %v\n", results)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
