package main

import (
	"fmt"
	"math"
)

func minFlipsMonoIncr(S string) int {
	n := len(S)
	onesToRight := make([]int, n)
	onesToLeft := make([]int, n)

	for i := len(onesToRight) - 1; i >= 0; i-- {
		if i == len(onesToRight)-1 {
			onesToRight[i] = 0

			continue
		}

		if S[i+1] == '1' {
			onesToRight[i] = onesToRight[i+1] + 1
		} else {
			onesToRight[i] = onesToRight[i+1]
		}
	}

	for i := 0; i < len(onesToLeft); i++ {
		if i == 0 {
			onesToLeft[i] = 0

			continue
		}

		if S[i-1] == '1' {
			onesToLeft[i] = onesToLeft[i-1] + 1
		} else {
			onesToLeft[i] = onesToLeft[i-1]
		}
	}

	minCost := math.MaxInt32

	for i := 0; i < n; i++ {
		numZerosToRight := (n - i - 1) - onesToRight[i]
		convCost := onesToLeft[i] + numZerosToRight

		if convCost < minCost {
			minCost = convCost
		}
	}

	return minCost
}

func main() {

	inputs := []string{
		"00110",
		"010110",
		"00011000",
		"1",
		"0",
		"00",
		"11",
	}

	for _, input := range inputs {
		fmt.Printf("convCost(%s) = %d\n", input, minFlipsMonoIncr(input))
	}

}
