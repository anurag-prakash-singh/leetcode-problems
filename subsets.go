package main

import "fmt"

func computePowerSet(nums []int, pos int, setSoFar []int, results [][]int) [][]int {
	if pos >= len(nums) {
		resultSet := make([]int, len(setSoFar))
		copy(resultSet, setSoFar)
		return append(results, resultSet)
	}

	// Case 1: move forward, excluding nums[pos]
	tempSetWithoutPos := make([]int, len(setSoFar))
	copy(tempSetWithoutPos, setSoFar)

	newResults := computePowerSet(nums, pos+1, tempSetWithoutPos, results)

	// Case 2: move forwawrd, including nums[pos]
	tempSetWithPos := []int{}
	tempSetWithPos = append(tempSetWithPos, setSoFar...)
	tempSetWithPos = append(tempSetWithPos, nums[pos])

	newResults = computePowerSet(nums, pos+1, tempSetWithPos, newResults)

	return newResults
}

func subsets(nums []int) [][]int {
	return computePowerSet(nums, 0, []int{}, [][]int{})
}

func test1() {
	results := subsets([]int{1, 2, 3})

	fmt.Printf("power set: %v\n", results)
}

func test2() {
	results := subsets([]int{})

	fmt.Printf("power set: %v\n", results)
}

func test3() {
	results := subsets([]int{1})

	fmt.Printf("power set: %v\n", results)
}

func main() {
	test1()
	test2()
	test3()
}
