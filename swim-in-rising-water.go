package main

import (
	"fmt"
)

func dfs(grid [][]int, visited map[int]bool, startRow, startCol, time int) {
	n := len(grid)
	deltaR := []int{-1, 1, 0, 0}
	deltaC := []int{0, 0, -1, 1}

	if visited[grid[startRow][startCol]] || grid[startRow][startCol] > time {
		return
	}

	visited[grid[startRow][startCol]] = true

	for d := 0; d < len(deltaR); d++ {
		nextRow, nextColumn := startRow+deltaR[d], startCol+deltaC[d]

		if nextRow < 0 || nextRow >= n || nextColumn < 0 || nextColumn >= n {
			continue
		}

		dfs(grid, visited, nextRow, nextColumn, time)
	}
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	time := 0

	for ; time < n*n; time++ {
		visited := make(map[int]bool)
		dfs(grid, visited, 0, 0, time)

		if visited[grid[n-1][n-1]] {
			break
		}
	}

	return time
}

func swimInWaterBinSearch(grid [][]int) int {
	n := len(grid)
	low := 0
	high := n*n - 1
	timesTried := make(map[int]bool)

	visited := make(map[int]bool)
	dfs(grid, visited, 0, 0, low)
	timesTried[low] = visited[grid[n-1][n-1]]

	visited = make(map[int]bool)
	dfs(grid, visited, 0, 0, high)
	timesTried[high] = visited[grid[n-1][n-1]]

	for true {
		mid := (low + high) / 2
		visited = make(map[int]bool)
		dfs(grid, visited, 0, 0, mid)
		timesTried[mid] = visited[grid[n-1][n-1]]

		if timesTried[mid] {
			if mid == 0 {
				return mid
			}

			if b, o := timesTried[mid-1]; o {
				if !b {
					return mid
				}
			}

			high = mid
		} else {
			low = mid + 1
		}
	}

	// Control should never come here
	return n*n - 1
}

func test1() {
	grid := [][]int{
		[]int{0, 2},
		[]int{1, 3},
	}

	result := swimInWater(grid)

	fmt.Printf("result: %d\n", result)
}

func test2() {
	grid := [][]int{
		[]int{0, 1, 2, 3, 4},
		[]int{24, 23, 22, 21, 5},
		[]int{12, 13, 14, 15, 16},
		[]int{11, 17, 18, 19, 20},
		[]int{10, 9, 8, 7, 6},
	}

	result := swimInWaterBinSearch(grid)

	fmt.Printf("result: %d\n", result)
}

func main() {
	test1()
	test2()
}
