package main

import (
	"fmt"
)

func dumpNonZeroCounts(counts []int) {
	for i, count := range counts {
		if count == 0 {
			continue
		}

		fmt.Printf("%c: %d; ", rune(i), count)
	}

	fmt.Println()
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	pChars := []rune(p)
	sChars := []rune(s)
	pCharCounts := make([]int, 256)
	sWindowCharCounts := make([]int, 256)

	for _, pChar := range pChars {
		pCharCounts[pChar] += 1
	}

	for i := 0; i < len(pChars); i++ {
		sWindowCharCounts[sChars[i]] += 1
	}

	numDiffs := 0

	for i := 0; i < len(pCharCounts); i++ {
		if pCharCounts[i] != sWindowCharCounts[i] {
			numDiffs++
		}
	}

	start, end := 0, len(pChars)-1
	anagramOffsets := make([]int, 0, len(pChars))

	for end < len(sChars) {
		if numDiffs == 0 {
			anagramOffsets = append(anagramOffsets, start)
		}

		sWindowCharCounts[sChars[start]]--

		if sWindowCharCounts[sChars[start]] == pCharCounts[sChars[start]] {
			numDiffs--
		} else if sWindowCharCounts[sChars[start]]+1 == pCharCounts[sChars[start]] {
			numDiffs++
		}

		start++

		end++

		if end >= len(sChars) {
			break
		}

		sWindowCharCounts[sChars[end]]++

		if sWindowCharCounts[sChars[end]] == pCharCounts[sChars[end]] {
			numDiffs--
		} else if sWindowCharCounts[sChars[end]]-1 == pCharCounts[sChars[end]] {
			numDiffs++
		}

		// fmt.Printf("At start=%d: ", start)
		// dumpNonZeroCounts(sWindowCharCounts)
	}

	return anagramOffsets
}

type testcase struct {
	s              string
	p              string
	expectedResult []int
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, n := range a {
		if n != b[i] {
			return false
		}
	}

	return true
}

func runTests() {
	testcases := []testcase{
		testcase{"cbaebabacd", "abc", []int{0, 6}},
		testcase{"abab", "ab", []int{0, 1, 2}},
	}

	for i, tc := range testcases {
		result := findAnagrams(tc.s, tc.p)

		if !slicesEqual(result, tc.expectedResult) {
			fmt.Printf("Test %d failed; expected: %v; got %v\n", i, tc.expectedResult, result)
		}
	}
}

func main() {
	runTests()
}
