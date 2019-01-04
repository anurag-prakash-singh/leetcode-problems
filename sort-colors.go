package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func sortColors(nums []int) {
	z := 0
	t := len(nums) - 1
	k := 0
	n := len(nums)

	for k <= t {
		// move z point to the first element that's not 0
		for z < n && nums[z] == 0 {
			z++
			k = max(k, z)
		}

		if z >= n {
			return
		}

		for t >= 0 && nums[t] == 2 {
			t--
		}

		if t < 0 || k > t {
			return
		}

		if nums[k] == 0 {
			nums[k], nums[z] = nums[z], nums[k]
			continue
		}

		if nums[k] == 2 {
			nums[k], nums[t] = nums[t], nums[k]
			continue
		}

		k++
	}
}

func test1() {
	nums := []int{2, 1, 0}

	sortColors(nums)

	fmt.Printf("sorted: %v\n", nums)
}

func test2() {
	nums := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums)

	fmt.Printf("sorted: %v\n", nums)
}

func test3() {
	nums := []int{1, 1, 1, 0, 0, 2, 1, 0, 0, 1, 2}

	sortColors(nums)

	fmt.Printf("sorted: %v\n", nums)
}

func test4() {
	nums := []int{0, 1, 2}

	sortColors(nums)

	fmt.Printf("sorted: %v\n", nums)
}

func main() {
	// test1()
	// test2()
	test3()
	// test4()
}
