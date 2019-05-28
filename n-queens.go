package main

import (
	"fmt"
	"strings"
)

func collisions(board [][]rune, r, c int) bool {
	n := len(board)
	qs := 0

	// check row
	for col := 0; col < n; col++ {
		if board[r][col] == 'Q' {
			qs++
		}
	}

	if qs > 1 {
		return true
	}

	qs = 0

	// check column
	for row := 0; row < n; row++ {
		if board[row][c] == 'Q' {
			qs++
		}
	}

	if qs > 1 {
		return true
	}

	// check diagonals
	dr := []int{-1 /*top-right*/, -1 /*top-left*/, 1 /*bottom-right*/, 1 /*bottom-left*/}
	dc := []int{1, -1, 1, -1}

	for dir := 0; dir < len(dr); dir++ {
		qs = 0

		for newR, newC := r, c; newR >= 0 && newR < n && newC >= 0 && newC < n; newR, newC = newR+dr[dir], newC+dc[dir] {
			if board[newR][newC] == 'Q' {
				qs++
			}
		}

		if qs > 1 {
			return true
		}
	}

	// No collisions detected
	return false
}

func placeQueens(board [][]rune, row int, solutions [][]string) [][]string {
	n := len(board)

	if row == n {
		newSolution := []string{}
		for _, boardRow := range board {
			newSolution = append(newSolution, string(boardRow))
		}

		return append(solutions, newSolution)
	}

	for c := 0; c < n; c++ {
		board[row][c] = 'Q'
		if !collisions(board, row, c) {
			// fmt.Printf("No col at %d, %d\n", row, c)
			solutions = placeQueens(board, row+1, solutions)
		} else {
			// fmt.Printf("%v\n", board)
		}
		board[row][c] = '.'
	}

	return solutions
}

func solveNQueens(n int) [][]string {
	solutions := [][]string{}

	if n == 0 {
		return solutions
	}

	board := make([][]rune, n)

	for i, _ := range board {
		board[i] = []rune(strings.Repeat(".", n))
	}

	return placeQueens(board, 0, solutions)
}

func test() {
	n := 4
	solutions := solveNQueens(n)

	fmt.Printf("solutions: %v\n", solutions)
}

func testCollisions() {
	// board := [][]rune{
	// 	[]rune(".Q.Q"),
	// 	[]rune("...."),
	// 	[]rune(".Q.."),
	// 	[]rune("...."),
	// }

	board1 := [][]rune{
		[]rune("Q"),
	}

	board2 := [][]rune{
		[]rune("Q."),
		[]rune(".Q"),
	}

	// fmt.Printf("collisions: %v\n", collisions(board, 0, 3))
	// fmt.Printf("collisions: %v\n", collisions(board, 2, 1))
	// fmt.Printf("collisions: %v\n", collisions(board, 0, 1))
	// fmt.Printf("collisions: %v\n", collisions(board, 2, 0))

	fmt.Printf("collisions: %v\n", collisions(board1, 0, 0))

	fmt.Printf("collisions: %v\n", collisions(board2, 0, 0))
	fmt.Printf("collisions: %v\n", collisions(board2, 1, 1))
}

func main() {
	test()
	// testCollisions()
}
