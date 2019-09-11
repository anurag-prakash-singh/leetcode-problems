package main

import (
	"fmt"
)

const (
	UnvisitedLand = '1'
	VisitedLand = '2'
)

func dfs(grid [][]byte, atR, atC int) {
	if grid[atR][atC] != UnvisitedLand {
		return
	}

	grid[atR][atC] = VisitedLand

	numRows := len(grid)
	numCols := len(grid[0])
	dirR := []int{-1, 1, 0, 0}
	dirC := []int{ 0, 0,-1, 1}

	for i := 0; i < len(dirR); i++ {
		newR, newC := atR + dirR[i], atC + dirC[i]

		if newR < 0 || newR >= numRows || newC < 0 || newC >= numCols {
			continue
		}

		dfs(grid, newR, newC)
	}
}

func numIslands(grid [][]byte) int {
	n := 0
	numRows := len(grid)

	if numRows == 0 {
		return 0
	}

	numCols := len(grid[0])

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if grid[i][j] == UnvisitedLand {
				n++
				dfs(grid, i, j)
			}
		}
	}

	return n
}

func test1() {
	grid := [][]byte {
		[]byte {'1','1','1','1','0'},
		[]byte {'1','1','0','1','0'},
		[]byte {'1','1','0','0','0'},
		[]byte {'0','0','0','0','0'},
	}

	fmt.Printf("numIslands: %d\n", numIslands(grid))
}



func main() {
	test1()
}
