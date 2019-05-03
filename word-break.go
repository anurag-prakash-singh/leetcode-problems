package main

import "fmt"

/*
 * https://leetcode.com/problems/word-break/
 */

func wordBreakIneff(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)

	for _, word := range wordDict {
		wordDictMap[word] = true
	}

	sChars := []rune(s)
	n := len(sChars)
	wordBreakExists := make([][]bool, n)

	for i := 0; i < n; i++ {
		wordBreakExists[i] = make([]bool, n)
	}

	for l := 1; l <= n; l++ {
		for i := 0; i < n; i++ {
			// end offset (inclusive)
			e := i + l - 1

			if e >= n {
				break
			}

			checkWord := string(sChars[i : e+1])

			if wordDictMap[checkWord] {
				wordBreakExists[i][e] = true

				continue
			} else {
				// the word itself doesn't exist in dict. Check if the word is made of smaller, valid words.
				for j := i; j < e; j++ {
					if wordBreakExists[i][j] && wordBreakExists[j+1][e] {
						wordBreakExists[i][e] = true

						break
					}
				}
			}
		}
	}

	return wordBreakExists[0][n-1]
}

func checkInArr(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}

	return false
}

func wordBreak(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)

	for _, word := range wordDict {
		wordDictMap[word] = true
	}

	sChars := []rune(s)
	n := len(sChars)
	// wordkBreakExists tracks whether the substring s[0:x] can be broken into valid words or not
	wordBreakExists := make([]bool, n+1)

	wordBreakExists[0] = true

	for l := 1; l <= n; l++ {
		for j := 0; j < l; j++ {
			remainingWord := string(sChars[j:l])

			if wordBreakExists[j] && wordDictMap[remainingWord]) {
				wordBreakExists[l] = true

				break
			}
		}
	}

	return wordBreakExists[n]
}

type testCase struct {
	s              string
	dict           []string
	expectedResult bool
}

func tests() {
	testCases := []testCase{
		testCase{"leetcode", []string{"leet", "code"}, true},
		testCase{"leetsscode", []string{"leet", "leets", "sscode"}, true},
		testCase{"applepenapple", []string{"apple", "pen"}, true},
		testCase{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for i, tc := range testCases {
		result := wordBreak(tc.s, tc.dict)

		if result == tc.expectedResult {
			fmt.Printf("Testcase %d PASSED\n", i)
		} else {
			fmt.Printf("Testcase %d FAILED(result: %v, expectedResult: %v)\n", i, result, tc.expectedResult)
		}
	}
}

func main() {
	tests()
}
