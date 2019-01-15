package main

import "fmt"

func search(board [][]byte, visited [][]bool, word []byte, wordOffset int, r int, c int) bool {
	if wordOffset >= len(word) {
		return true
	}

	if board[r][c] != word[wordOffset] {
		return false
	}

	rows := len(board)
	cols := len(board[0])

	dirRow := []int{-1, 1, 0, 0}
	dirCol := []int{0, 0, -1, 1}

	otherCoordsChecked := false

	for i := 0; i < len(dirRow); i++ {
		nextRow, nextCol := r+dirRow[i], c+dirCol[i]

		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
			continue
		}

		if visited[nextRow][nextCol] {
			continue
		}

		otherCoordsChecked = true
		visited[nextRow][nextCol] = true
		exists := search(board, visited, word, wordOffset+1, nextRow, nextCol)
		visited[nextRow][nextCol] = false

		if exists {
			return true
		}
	}

	// To account for the case where there are no more unvisited directions
	// left to take but we've also reached the last letter of the word to check for (and it matched).
	if !otherCoordsChecked && wordOffset == (len(word)-1) {
		return true
	}

	return false
}

func initVisited(visited [][]bool) {
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[0]); j++ {
			visited[i][j] = false
		}
	}
}

func exist(board [][]byte, word string) bool {
	wordBytes := []byte(word)
	rows := len(board)

	if len(word) == 0 {
		return true
	}

	if rows == 0 {
		return false
	}

	cols := len(board[0])

	if cols == 0 {
		return false
	}

	visited := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			initVisited(visited)

			visited[r][c] = true
			if search(board, visited, wordBytes, 0, r, c) {
				return true
			}
		}
	}

	return false
}

func test1() {
	fmt.Printf("test1\n")

	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}

	fmt.Printf("ans: %v\n", exist(board, "ABCCED"))
	fmt.Printf("ans: %v\n", exist(board, "SEE"))
	fmt.Printf("ans: %v\n", exist(board, "ABCB"))
	fmt.Printf("ans: %v\n", exist(board, ""))
}

func test2() {
	fmt.Printf("test2\n")

	board := [][]byte{}

	fmt.Printf("ans: %v\n", exist(board, "SEE"))
	fmt.Printf("ans: %v\n", exist(board, ""))
}

func test3() {
	fmt.Printf("test3\n")

	board := [][]byte{}

	fmt.Printf("ans: %v\n", exist(board, "SEE"))
	fmt.Printf("ans: %v\n", exist(board, ""))
}

func test4() {
	fmt.Printf("test4\n")

	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'E', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}

	fmt.Printf("ans: %v\n", exist(board, "ABCESEEEFS"))
}

func test5() {
	fmt.Printf("test5\n")

	board := [][]byte{
		[]byte{'a', 'a'},
	}

	fmt.Printf("ans: %v\n", exist(board, "aaa"))
}

func main() {
	test1()
	// test2()
	// test4()
	test5()
}
