package main

import (
	"fmt"
	"math"
)

func minWindowUniqueTChars(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	tLUT := make([]int, 256)
	tBytes := []rune(t)

	for _, c := range tBytes {
		tLUT[byte(c)] = 1
	}

	sChars := []rune(s)
	windowCharCount := make([]int, 256)
	charPresence := make([]int, 256)
	charPresenceCount := 0

	start, end := 0, 0
	minWindowSize := math.MaxInt32
	minWindowStart, minWindowEnd := -1, -1

	if tLUT[int(sChars[start])] == 1 {
		windowCharCount[int(sChars[start])]++
		charPresence[int(sChars[start])] = 1
		charPresenceCount++

		if charPresenceCount == len(t) {
			if minWindowSize > (end - start + 1) {
				minWindowSize = end - start + 1
				minWindowStart, minWindowEnd = start, end
			}
		}
	}

	for end < len(s) {
		if charPresenceCount < len(t) {
			end++

			if end >= len(s) {
				break
			}

			// Add the character at end to our presence counts
			if tLUT[int(sChars[end])] == 1 {
				windowCharCount[int(sChars[end])]++

				if windowCharCount[int(sChars[end])] == 1 {
					charPresence[int(sChars[end])] = 1
					charPresenceCount++
				}
			}
		} else {
			start++

			if start > 0 && tLUT[int(sChars[start-1])] == 1 {
				windowCharCount[int(sChars[start-1])]--

				if windowCharCount[int(sChars[start-1])] == 0 {
					charPresence[int(sChars[start-1])] = 0
					charPresenceCount--
				}
			}

			if start > end {
				end = start

				if end >= len(s) {
					break
				}

				if tLUT[int(sChars[end])] == 1 {
					windowCharCount[int(sChars[end])]++

					if windowCharCount[int(sChars[end])] == 1 {
						charPresence[int(sChars[end])] = 1
						charPresenceCount++
					}
				}
			}
		}

		if charPresenceCount == len(t) {
			if minWindowSize > (end - start + 1) {
				minWindowSize = end - start + 1
				minWindowStart, minWindowEnd = start, end
			}
		}
	}

	if minWindowSize == 0 || minWindowSize == math.MaxInt32 {
		return ""
	}

	return string(sChars[minWindowStart : minWindowEnd+1])
}

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	tLUT := make([]int, 256)
	uniqueLUT := make([]int, 256)
	uniqueTChars := 0
	tBytes := []rune(t)

	for _, c := range tBytes {
		tLUT[int(c)]++

		if uniqueLUT[int(c)] == 0 {
			uniqueTChars++
		}

		uniqueLUT[int(c)] = 1
	}

	sChars := []rune(s)
	windowCharCount := make([]int, 256)
	charPresenceCount := 0

	start, end := 0, 0
	minWindowSize := math.MaxInt32
	minWindowStart, minWindowEnd := -1, -1

	if tLUT[int(sChars[start])] > 0 {
		windowCharCount[int(sChars[start])]++

		if windowCharCount[int(sChars[start])] == tLUT[int(sChars[start])] {
			charPresenceCount++
		}

		if charPresenceCount == uniqueTChars {
			if minWindowSize > (end - start + 1) {
				minWindowSize = end - start + 1
				minWindowStart, minWindowEnd = start, end
			}
		}
	}

	for end < len(s) {
		if charPresenceCount < uniqueTChars {
			end++

			if end >= len(s) {
				break
			}

			// Add the character at end to our presence counts
			if tLUT[int(sChars[end])] > 0 {
				windowCharCount[int(sChars[end])]++

				if windowCharCount[int(sChars[end])] == tLUT[int(sChars[end])] {
					charPresenceCount++
				}
			}
		} else {
			start++

			if start > 0 && tLUT[int(sChars[start-1])] > 0 {
				windowCharCount[int(sChars[start-1])]--

				if windowCharCount[int(sChars[start-1])] == tLUT[int(sChars[start-1])]-1 {
					charPresenceCount--
				}
			}

			if start > end {
				end = start

				if end >= len(s) {
					break
				}

				if tLUT[int(sChars[end])] > 0 {
					windowCharCount[int(sChars[end])]++

					if windowCharCount[int(sChars[end])] == tLUT[int(sChars[end])] {
						charPresenceCount++
					}
				}
			}
		}

		if charPresenceCount == uniqueTChars {
			if minWindowSize > (end - start + 1) {
				minWindowSize = end - start + 1
				minWindowStart, minWindowEnd = start, end
			}
		}
	}

	if minWindowSize == 0 || minWindowSize == math.MaxInt32 {
		return ""
	}

	return string(sChars[minWindowStart : minWindowEnd+1])
}

type testcase struct {
	input    string
	search   string
	expected string
}

func runTests() {
	testcases := []testcase{
		testcase{input: "ADOBECODEBANC", search: "ABC", expected: "BANC"},
		testcase{input: "ABC", search: "ABC", expected: "ABC"},
		testcase{input: "", search: "ABC", expected: ""},
		testcase{input: "A", search: "ABC", expected: ""},
		testcase{input: "AC", search: "ABC", expected: ""},
		testcase{input: "ACB", search: "ABC", expected: "ACB"},
		testcase{input: "ACB", search: "", expected: ""},
		testcase{input: "OOADCOBOOO", search: "ABC", expected: "ADCOB"},
		testcase{input: "AOOOOOBC", search: "ABC", expected: "AOOOOOBC"},
		testcase{input: "OAOOOOOBC", search: "ABC", expected: "AOOOOOBC"},
		testcase{input: "AA", search: "AA", expected: "AA"},
		testcase{input: "AA", search: "AAA", expected: ""},
	}

	for i, tc := range testcases {
		result := minWindow(tc.input, tc.search)

		if result != tc.expected {
			fmt.Printf("TC %d (%s, %s) FAILED. Expected: %s; got: %s\n", i, tc.input, tc.search, tc.expected, result)
		}
	}
}

func main() {
	runTests()
}
