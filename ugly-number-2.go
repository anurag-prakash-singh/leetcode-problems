package main

import "fmt"

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}

		return c
	}

	if b < c {
		return b
	}

	return c
}

func nthUglyNumber(n int) int {
	nums := make([]int, n)
	x, y, z := 0, 0, 0
	nums[0] = 1
	next := 1

	for next < len(nums) {
		a := 2 * nums[x]
		b := 3 * nums[y]
		c := 5 * nums[z]

		min := min3(a, b, c)
		nums[next] = min
		next++

		if a == min {
			x++
		}

		if b == min {
			y++
		}

		if c == min {
			z++
		}
	}

	return nums[len(nums)-1]
}

func test() {
	for i := 1; i < 11; i++ {
		uglyNumber := nthUglyNumber(i)

		fmt.Printf("%d ", uglyNumber)
	}

	fmt.Printf("\n")
}

func main() {
	test()
}
