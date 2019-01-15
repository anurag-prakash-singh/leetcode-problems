package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lastVal := nums[0]
	lastValCount := 0
	curr := 0
	la := 0

	for ; la < len(nums); la++ {
		if lastVal == nums[la] {
			lastValCount++
		} else {
			lastVal = nums[la]
			lastValCount = 1
		}

		if lastValCount < 3 {
			nums[curr] = nums[la]
			curr++
		}
	}

	return curr
}

func test1() {
	nums := []int{1, 1, 2, 3, 4}
	newLen := removeDuplicates(nums)

	fmt.Printf("ans = %v\n", nums[:newLen])
}

func test2() {
	nums := []int{1, 1, 1, 2, 3, 4}
	newLen := removeDuplicates(nums)

	fmt.Printf("ans = %v\n", nums[:newLen])
}

func test3() {
	nums := []int{1, 1, 1, 1}
	newLen := removeDuplicates(nums)

	fmt.Printf("ans = %v\n", nums[:newLen])
}

func test4() {
	nums := []int{1, 1, 1, 2, 2, 2, 2, 3, 3}
	newLen := removeDuplicates(nums)

	fmt.Printf("ans = %v\n", nums[:newLen])
}

func test5() {
	nums := []int{1, 2, 2, 2, 3}
	newLen := removeDuplicates(nums)

	fmt.Printf("ans = %v\n", nums[:newLen])
}

func test6() {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5, 5}
	newLen := removeDuplicates(nums)

	fmt.Printf("ans = %v\n", nums[:newLen])
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	test6()
}
