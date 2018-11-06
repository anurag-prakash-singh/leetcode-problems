package main

import (
	"fmt"
)

func threeEqualParts(A []int) []int {
	numOnes := 0

	for _, num := range A {
		if num == 1 {
			numOnes++
		}
	}

	if numOnes%3 != 0 {
		return []int{-1, -1}
	}

	if numOnes == 0 {
		return []int{0, len(A) - 1}
	}

	// Identify the potentially repeated binary value's string
	onesInValue := numOnes / 3
	onesSoFar := 0
	j := len(A) - 1

	for ; j >= 0; j-- {
		if A[j] == 1 {
			onesSoFar++
		}

		if onesSoFar == onesInValue {
			break
		}
	}

	// j should now point at the MSB of the repeated value
	// Look for it in the beginning
	i := 0

	for ; i < j && A[i] != 1; i++ {
	}

	if i == j {
		return []int{-1, -1}
	}

	for jj := j; jj < len(A); jj, i = jj+1, i+1 {
		if A[i] != A[jj] || i >= j {
			return []int{-1, -1}
		}
	}

	i--

	// We should now have the correct value of i
	k := i + 1

	for ; k < j && A[k] != 1; k++ {
	}

	for jj := j; jj < len(A); jj, k = jj+1, k+1 {
		if A[k] != A[jj] || k >= j {
			return []int{-1, -1}
		}
	}

	// Probably a superfluous check
	for t := k + 1; t < j; t++ {
		if A[t] == 1 {
			return []int{-1, -1}
		}
	}

	j = k

	return []int{i, j}
}

func main() {

	ans1 := threeEqualParts([]int{1, 1, 0, 1, 1})
	ans2 := threeEqualParts([]int{1, 1, 1})
	ans3 := threeEqualParts([]int{1, 0, 1, 0, 1})

	fmt.Printf("ans1: %v\n", ans1)
	fmt.Printf("ans2: %v\n", ans2)
	fmt.Printf("ans3: %v\n", ans3)

}
