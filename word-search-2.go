package main

import "fmt"

func runeToByteBoard(rBoard [][]rune) [][]byte {
	rows := len(rBoard)
	cols := len(rBoard[0])
	bBoard := make([][]byte, rows)

	for r := 0; r < rows; r++ {
		bBoard[r] = make([]byte, cols)

		for c := 0; c < cols; c++ {
			bBoard[r][c] = byte(rBoard[r][c])
		}
	}

	return bBoard
}

func clearTraversed(traversed [][]bool) {
	for r := 0; r < len(traversed); r++ {
		for c := 0; c < len(traversed[0]); c++ {
			traversed[r][c] = false
		}
	}
}

func checkTraversed(traversed [][]bool) {
	for r := 0; r < len(traversed); r++ {
		for c := 0; c < len(traversed[0]); c++ {
			if traversed[r][c] {
				fmt.Printf("traversed[%d][%d] is true", r, c)

				return
			}
		}
	}
}

func checkForWord(board [][]byte, traversed [][]bool, word []rune, fromRow, fromCol, start int) bool {
	if start >= len(word) {
		return true
	}

	if board[fromRow][fromCol] != byte(word[start]) {
		return false
	}

	if start == len(word)-1 {
		return true
	}

	dirR := []int{0, 0, -1, 1}
	dirC := []int{-1, 1, 0, 0}
	numRows := len(board)
	numCols := len(board[0])

	for d := 0; d < len(dirR); d++ {
		newR, newC := fromRow+dirR[d], fromCol+dirC[d]

		if newR < 0 || newR >= numRows || newC < 0 || newC >= numCols || traversed[newR][newC] {
			continue
		}

		traversed[newR][newC] = true

		if checkForWord(board, traversed, word, newR, newC, start+1) {
			traversed[newR][newC] = false
			return true
		}

		traversed[newR][newC] = false
	}

	return false
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 {
		return []string{}
	}

	rows := len(board)
	cols := len(board[0])
	result := []string{}

	traversed := make([][]bool, rows)

	for r := 0; r < rows; r++ {
		traversed[r] = make([]bool, cols)
	}

	for w := 0; w < len(words); w++ {
		wordRunes := []rune(words[w])
		wordFound := false

		for r := 0; r < rows && !wordFound; r++ {
			for c := 0; c < cols && !wordFound; c++ {
				// checkTraversed(traversed)

				if (byte)(wordRunes[0]) != board[r][c] {
					continue
				}

				// fmt.Printf("found ")
				traversed[r][c] = true
				if checkForWord(board, traversed, wordRunes, r, c, 0) {
					traversed[r][c] = false
					wordFound = true
					break
				}
				traversed[r][c] = false
			}
		}

		if wordFound {
			result = append(result, words[w])
		}
	}

	return result
}

func test1() {
	rBoard := [][]rune{
		[]rune{'o', 'a', 'a', 'n'},
		[]rune{'e', 't', 'a', 'e'},
		[]rune{'i', 'h', 'k', 'r'},
		[]rune{'i', 'f', 'l', 'v'},
	}
	bBoard := runeToByteBoard(rBoard)
	words := []string{"oath", "pea", "eat", "rain"}
	result := findWords(bBoard, words)

	fmt.Printf("result: %v\n", result)
}

func main() {
	test1()
}
