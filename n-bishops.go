package main

import (
	"bufio"
	"fmt"
	"os"
)

type bishopPos struct {
	row int
	col int
}

type board struct {
	bishops []bishopPos
	maxRows int
	maxCols int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func readBishopPos() board {
	scanner := bufio.NewScanner(os.Stdin)
	maxRows, maxCols := -1, -1

	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)
	bps := make([]bishopPos, n)

	// fmt.Printf("Reading %d lines\n", n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &(bps[i].row), &(bps[i].col))
		maxCols = max(maxCols, bps[i].col)
	}

	for _, bp := range bps {
		diagL := bp.row + bp.col - 1
		diagR := bp.row + (maxCols - bp.col)

		maxRows = max(max(maxRows, diagL), diagR)
	}

	return board{
		bishops: bps,
		maxRows: maxRows,
		maxCols: maxCols,
	}
}

func countCaptures(b board) int {
	diagLCounts := make([]int, b.maxRows+1)
	diagRCounts := make([]int, b.maxRows+1)
	captureCount := 0

	for _, bp := range b.bishops {
		diagL := bp.row + bp.col - 1
		diagR := bp.row + (b.maxCols - bp.col)

		prevDiagLCount := diagLCounts[diagL]
		prevDiagRCount := diagRCounts[diagR]

		diagLCounts[diagL]++
		diagRCounts[diagR]++

		prevDiagLCaptures := prevDiagLCount * (prevDiagLCount - 1) / 2
		prevDiagRCaptures := prevDiagRCount * (prevDiagRCount - 1) / 2
		newDiagLCaptures := diagLCounts[diagL] * (diagLCounts[diagL] - 1) / 2
		newDiagRCaptures := diagRCounts[diagR] * (diagRCounts[diagR] - 1) / 2

		captureCount = captureCount + (newDiagLCaptures - prevDiagLCaptures) + (newDiagRCaptures - prevDiagRCaptures)
	}

	return captureCount
}

func main() {
	b := readBishopPos()

	// fmt.Printf("%v\n", b)
	fmt.Printf("%d\n", countCaptures(b))
}
