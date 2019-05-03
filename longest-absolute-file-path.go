package main

import (
	"fmt"
	"strings"
)

func indexOfRune(chars []rune, char rune) int {
	for i, c := range chars {
		if c == char {
			return i
		}
	}

	return -1
}

func nextToken(chars []rune) string {
	if len(chars) == 0 {
		return ""
	}

	if chars[0] == '\\' {
		wspChars := []rune{'\\', chars[1]}

		return string(wspChars)
	}

	if chars[0] == '\n' || chars[0] == '\t' {
		wspChars := []rune{'\\', chars[1]}

		return string(wspChars)
	}

	slashPos := indexOfRune(chars, '\\')

	if slashPos == -1 {
		return string(chars)
	}

	return string(chars[:slashPos])
}

func nestingLevel(str string) int {
	chars := []rune(str)

	for i, c := range chars {
		if c != '\t' {
			return i
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthLongestPath(input string) int {
	lines := strings.Split(input, "\n")
	dirStack := []string{}
	dirLen := 0
	maxPathLen := 0

	for _, line := range lines {
		nstLvl := nestingLevel(line)

		for nstLvl < len(dirStack) {
			dirLen -= len(dirStack[len(dirStack)-1])
			dirStack = dirStack[:len(dirStack)-1]
		}

		trimmed := strings.Trim(line, "\t")

		if strings.Index(line, ".") != -1 {
			// file
			// Add len(dirStack) to account for the file component separator ('/')
			pathLen := dirLen + len(dirStack) + len(trimmed)
			maxPathLen = max(maxPathLen, pathLen)
		} else {
			// dir
			dirLen += len(trimmed)
			dirStack = append(dirStack, trimmed)
		}
	}

	return maxPathLen
}

func testNextToken() {
	tests := []string{
		"abcd",
		"a\\nb\\n\\tc",
		"dir\\n\\tsubdir1\\n\\tsubdir2\\n\\t\\tfile.ext",
		"a\nb\n\tc",
	}

	for _, test := range tests {
		testChars := []rune(test)

		for token := nextToken(testChars); token != ""; {
			fmt.Printf("token: %s\n", token)

			testChars = testChars[len(token):]
			token = nextToken(testChars)
		}
	}
}

type testCase struct {
	path              string
	longestPathLength int
}

func test() {
	testCases := []testCase{
		testCase{path: "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", longestPathLength: 20},
		testCase{path: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", longestPathLength: 32},
		testCase{path: "dir\n    file.txt", longestPathLength: 12},
	}

	for i, tc := range testCases {
		result := lengthLongestPath(tc.path)

		if result != tc.longestPathLength {
			fmt.Printf("Test number %d FAILED. Expected: %d; got: %d\n", i, tc.longestPathLength, result)
		} else {
			fmt.Printf("Test number %d PASSED\n", i)
		}
	}
}

func main() {
	test()
}
