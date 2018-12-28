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

func maxWidthRamp(A []int) int {
	rampSizeStack := []int{}

	for i, s := range A {
		if len(rampSizeStack) == 0 || A[rampSizeStack[len(rampSizeStack)-1]] > s {
			rampSizeStack = append(rampSizeStack, i)
		}
	}

	maxWidth := 0

	for i := len(A) - 1; i >= 0 && len(rampSizeStack) > 0; i-- {
		if rampSizeStack[len(rampSizeStack)-1] > i {
			rampSizeStack = rampSizeStack[:len(rampSizeStack)-1]

			continue
		}

		for len(rampSizeStack) > 0 && A[rampSizeStack[len(rampSizeStack)-1]] <= A[i] {
			maxWidth = max(maxWidth, i-rampSizeStack[len(rampSizeStack)-1])

			rampSizeStack = rampSizeStack[:len(rampSizeStack)-1]
		}
	}

	return maxWidth
}

func test1() {
	A := []int{6, 0, 8, 2, 1, 5}

	fmt.Printf("ans = %d\n", maxWidthRamp(A))
}

func test2() {
	A := []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}

	fmt.Printf("ans = %d\n", maxWidthRamp(A))
}

func test3() {
	A := []int{9, 8, 1}

	fmt.Printf("ans = %d\n", maxWidthRamp(A))
}

func main() {
	test1()
	test2()
	test3()
}
