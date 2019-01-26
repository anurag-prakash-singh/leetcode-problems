package main

import "fmt"

func countDecodings(nums []rune, remaining int, cache map[int]int) int {
	if d, ok := cache[remaining]; ok {
		return d
	}

	if remaining == 0 {
		return 1
	}

	if remaining < 0 {
		return 0
	}

	// Case 1: take 1
	num1 := nums[0] - '0'

	if num1 == 0 {
		return 0
	}

	decodings := countDecodings(nums[1:], remaining-1, cache)
	cache[remaining-1] = decodings

	// Case 2: take 2 if possible
	if len(nums) < 2 {
		return decodings
	}

	num2 := nums[1] - '0'
	num12 := 10*num1 + num2

	if num12 >= 10 && num12 <= 26 {
		decodings12 := countDecodings(nums[2:], remaining-2, cache)
		cache[remaining-2] = decodings12
		decodings += decodings12
	}

	return decodings
}

func numDecodings(s string) int {
	cache := make(map[int]int)
	return countDecodings([]rune(s), len(s), cache)
}

func test1() {
	fmt.Printf("ans = %d\n", numDecodings("12"))
}

func test2() {
	fmt.Printf("ans = %d\n", numDecodings("226"))
}

func main() {
	test1()
	test2()
}
