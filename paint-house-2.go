package main

import "fmt"

func min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Assumption: all numbers in arr are non-negative
func fillMinExceptAt(arr []int, minLeftScratch []int, minRightScratch []int) {
	// Populate minLeftScratch
	for i := range arr {
		if i == 0 {
			minLeftScratch[i] = -1

			continue
		}

		if i == 1 {
			minLeftScratch[i] = arr[i-1]

			continue
		}

		minLeftScratch[i] = min2(minLeftScratch[i-1], arr[i-1])
	}

	// Populate minRightScratch
	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			minRightScratch[i] = -1

			continue
		}

		if i == len(arr)-2 {
			minRightScratch[i] = arr[i+1]

			continue
		}

		minRightScratch[i] = min2(minRightScratch[i+1], arr[i+1])
	}

	for i := 0; i < len(arr); i++ {
		if minLeftScratch[i] == -1 {
			arr[i] = minRightScratch[i]

			continue
		}

		if minRightScratch[i] == -1 {
			arr[i] = minLeftScratch[i]

			continue
		}

		arr[i] = min2(minLeftScratch[i], minRightScratch[i])
	}
}

func calcMin(costs [][]int, minCosts [][]int) int {
	k := len(costs[0])
	n := len(costs)
	lScratch := make([]int, k)
	rScratch := make([]int, k)

	for i := n - 1; i >= 1; i-- {
		if i == n-1 {
			for j := 0; j < k; j++ {
				minCosts[i][j] = costs[i][j]
			}

			fillMinExceptAt(minCosts[i], lScratch, rScratch)

			continue
		}

		for j := 0; j < k; j++ {
			minCosts[i][j] = costs[i][j] + minCosts[i+1][j]
		}

		fillMinExceptAt(minCosts[i], lScratch, rScratch)
	}

	// Handle final case of the first house
	var resultMinCost int

	for j := 0; j < k; j++ {
		if n > 1 {
			minCosts[0][j] = costs[0][j] + minCosts[1][j]
		} else {
			minCosts[0][j] = costs[0][j]
		}

		if j == 0 {
			resultMinCost = minCosts[0][j]

			continue
		}

		resultMinCost = min2(resultMinCost, minCosts[0][j])
	}

	return resultMinCost
}

func minCostII(costs [][]int) int {
	n := len(costs)

	if n == 0 {
		return 0
	}

	k := len(costs[0])
	minCosts := make([][]int, n)

	for i := 0; i < n; i++ {
		minCosts[i] = make([]int, k)
	}

	return calcMin(costs, minCosts)
}

func test1() {
	test1 := [][]int{
		[]int{17, 2, 17},
		[]int{16, 16, 5},
		[]int{14, 3, 19},
	}

	result := minCostII(test1)

	fmt.Printf("result: %d\n", result)
}

func test2() {
	test1 := [][]int{
		[]int{17, 2, 17},
	}

	result := minCostII(test1)

	fmt.Printf("result: %d\n", result)
}

func test3() {
	test1 := [][]int{
		[]int{1, 5, 3},
		[]int{2, 9, 4},
	}

	result := minCostII(test1)

	fmt.Printf("result: %d\n", result)
}

func testFillMinExcept() {
	// nums := []int{17, 2, 17}
	nums := []int{17, 17}
	lScratch := make([]int, len(nums))
	rScratch := make([]int, len(nums))

	fillMinExceptAt(nums, lScratch, rScratch)

	fmt.Printf("minExceptAt: %v\n", nums)
}

func main() {
	test1()
	test2()
	test3()
	// testFillMinExcept()
}
