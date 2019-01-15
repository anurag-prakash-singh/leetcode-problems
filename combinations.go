package main

import (
	"fmt"
)

func pick(nn []int, k int, curr []int, results [][]int) [][]int {
	if k == 0 {
		if len(curr) > 0 {
			newResult := make([]int, len(curr))
			copy(newResult, curr)
			return append(results, newResult)
		}

		return results
	}

	if len(nn) == 0 {
		return results
	}

	for i := 0; i < len(nn); i++ {
		if len(curr) > 0 && nn[i] < curr[len(curr)-1] {
			continue
		}
		nnCopyNoI := make([]int, len(nn)-1)
		copy(nnCopyNoI[0:i], nn[0:i])
		copy(nnCopyNoI[i:], nn[i+1:])
		curr = append(curr, nn[i])
		results = pick(nnCopyNoI, k-1, curr, results)
		curr = curr[0 : len(curr)-1]
	}

	return results
}

func pickBitmask(nn []int, picked []bool, k int, curr []int, results [][]int) [][]int {
	if k == 0 {
		if len(curr) > 0 {
			newResult := make([]int, len(curr))
			copy(newResult, curr)
			return append(results, newResult)
		}

		return results
	}

	if len(nn) == 0 {
		return results
	}

	for i := 0; i < len(nn); i++ {
		if picked[i] {
			continue
		}

		if len(curr) > 0 && nn[i] < curr[len(curr)-1] {
			continue
		}

		picked[i] = true
		curr = append(curr, nn[i])
		results = pickBitmask(nn, picked, k-1, curr, results)
		curr = curr[0 : len(curr)-1]
		picked[i] = false
	}

	return results
}

func combine(n int, k int) [][]int {
	nn := make([]int, n)

	for i := 1; i <= n; i++ {
		nn[i-1] = i
	}

	results := make([][]int, 0)
	picked := make([]bool, n)

	return pickBitmask(nn, picked, k, []int{}, results)
}

func test1() {
	ans := combine(4, 2)

	fmt.Printf("num: %d; ans: %v\n", len(ans), ans)
}

func test2() {
	ans := combine(4, 3)

	fmt.Printf("num: %d; ans: %v\n", len(ans), ans)
}

func test3() {
	ans := combine(3, 3)

	fmt.Printf("num: %d; ans: %v\n", len(ans), ans)
}

func main() {
	test1()
	test2()
	test3()
}
