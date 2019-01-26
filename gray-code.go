package main

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	if n == 1 {
		return []int{0, 1}
	}

	codes := make([]int, 1<<uint(n))
	partialCodes := grayCode(n - 1)

	for i := 0; i < len(partialCodes); i++ {
		var first, second int

		if i%2 == 0 {
			first, second = partialCodes[i], (partialCodes[i] | (1 << uint(n-1)))
		} else {
			second, first = partialCodes[i], (partialCodes[i] | (1 << uint(n-1)))
		}

		codes[2*i] = first
		codes[2*i+1] = second
	}

	return codes
}

func test1() {
	fmt.Printf("%v\n", grayCode(2))
}

func test2() {
	fmt.Printf("%v\n", grayCode(3))
}

func main() {
	// test1()
	test2()
}
