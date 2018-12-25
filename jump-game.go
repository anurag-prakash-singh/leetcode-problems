package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func canJump(nums []int) bool {
	maxJumpableIndex := 0

	if len(nums) == 0 || len(nums) == 1 {
		return true
	}

	for i, jmpLen := range nums {
		if i > maxJumpableIndex {
			return false
		}

		maxJumpableIndex = max(maxJumpableIndex, i+jmpLen)
	}

	return true
}

func test1() {
	nums := []int{2, 3, 1, 1, 4}

	fmt.Printf("can jump? %v\n", canJump(nums))
}

func test2() {
	nums := []int{3, 2, 1, 0, 4}

	fmt.Printf("can jump? %v\n", canJump(nums))
}

func test3() {
	nums := []int{0}

	fmt.Printf("can jump? %v\n", canJump(nums))
}

func main() {
	test1()
	test2()
	test3()
}
