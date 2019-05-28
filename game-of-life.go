package main

import "fmt"

func getOrig(n int) int {
	if n < 10 {
		return n
	}

	if n == 10 {
		return 0
	}

	return 1
}

func getNew(n int) int {
	if n < 10 {
		return n
	}

	if n == 10 {
		return 1
	}

	return 0
}

func countLiveNeighbors(board [][]int, atR, atC int) int {
	dirR := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dirC := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	numRows := len(board)
	numCols := len(board[0])
	numLiveNeighbors := 0

	for i := 0; i < len(dirR); i++ {
		nbrR, nbrC := atR+dirR[i], atC+dirC[i]

		if nbrR < 0 || nbrR >= numRows || nbrC < 0 || nbrC >= numCols {
			continue
		}

		numLiveNeighbors += getOrig(board[nbrR][nbrC])
	}

	return numLiveNeighbors
}

func gameOfLife(board [][]int) {
	numRows := len(board)
	numCols := len(board[0])

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			numLiveNeighbors := countLiveNeighbors(board, i, j)

			if numLiveNeighbors < 2 {
				if board[i][j] == 1 {
					board[i][j] = 11
				}
			} else if numLiveNeighbors < 4 {
				// 2 or 3
				if numLiveNeighbors == 3 && getOrig(board[i][j]) == 0 {
					board[i][j] = 10
				}
			} else {
				// more than 3
				if getOrig(board[i][j]) == 1 {
					board[i][j] = 11
				}
			}
		}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			board[i][j] = getNew(board[i][j])
		}
	}
}

func test1() {
	board := [][]int{
		[]int{0, 1, 0},
		[]int{0, 0, 1},
		[]int{1, 1, 1},
		[]int{0, 0, 0},
	}

	gameOfLife(board)

	fmt.Printf("new board: %v\n", board)
}

func main() {
	test1()
}
