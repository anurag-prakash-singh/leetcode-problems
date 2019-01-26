package main

import "fmt"

func searchInRange(nums []int, low, high, target int) bool {
	if low > high {
		return false
	}

	mid := (low + high) / 2

	if target == nums[mid] {
		return true
	} else {
		return (searchInRange(nums, low, mid-1, target) || searchInRange(nums, mid+1, high, target))
	}
}

func search(nums []int, target int) bool {
	return searchInRange(nums, 0, len(nums)-1, target)
}

func test1() {
	nums := []int{3, 1}

	fmt.Printf("ans = %v\n", search(nums, 1))
}

func main() {
	test1()
}
