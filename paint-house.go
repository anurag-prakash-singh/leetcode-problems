package main

import "fmt"

const (
	RED   = 0
	GREEN = 1
	BLUE  = 2
)

func min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func min3(a, b, c int) int {
	return min2(a, min2(b, c))
}

func calcMinAt(at int, withColor int, costs [][]int, minCosts [][]int) {
	if at >= len(costs) {
		return
	}

	if at == len(costs)-1 {
		minCosts[at][RED] = costs[at][RED]
		minCosts[at][GREEN] = costs[at][GREEN]
		minCosts[at][BLUE] = costs[at][BLUE]

		return
	}

	switch withColor {
	case RED:
		minCosts[at][RED] = costs[at][RED] + min2(minCosts[at+1][GREEN], minCosts[at+1][BLUE])
	case BLUE:
		minCosts[at][BLUE] = costs[at][BLUE] + min2(minCosts[at+1][GREEN], minCosts[at+1][RED])
	case GREEN:
		minCosts[at][GREEN] = costs[at][GREEN] + min2(minCosts[at+1][BLUE], minCosts[at+1][RED])
	}
}

func minCost(costs [][]int) int {
	n := len(costs)

	if n == 0 {
		return 0
	}

	minCosts := make([][]int, n)

	for i := 0; i < n; i++ {
		minCosts[i] = make([]int, 3)
	}

	for i := n - 1; i >= 0; i-- {
		calcMinAt(i, RED, costs, minCosts)
		calcMinAt(i, GREEN, costs, minCosts)
		calcMinAt(i, BLUE, costs, minCosts)
	}

	result := min3(minCosts[0][RED], minCosts[0][BLUE], minCosts[0][GREEN])

	return result
}

func test1() {
	test1 := [][]int{
		[]int{17, 2, 17},
		[]int{16, 16, 5},
		[]int{14, 3, 19},
	}

	result := minCost(test1)

	fmt.Printf("result: %d\n", result)
}

func test2() {
	test1 := [][]int{
		[]int{17, 2, 17},
	}

	result := minCost(test1)

	fmt.Printf("result: %d\n", result)
}

func main() {
	test1()
	test2()
}
