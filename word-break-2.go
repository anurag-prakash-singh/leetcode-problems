package main

import "fmt"

/*
 * https://leetcode.com/problems/word-break-ii/
 */

func traverseBreaks(at int, sChars []rune, breaks [][]wordBreakNode, acc []string, strSoFar string) []string {
	if at == 0 {
		acc = append(acc, strSoFar)

		return acc
	}

	followingWordOffsets := breaks[at]

	for i := len(followingWordOffsets) - 1; i >= 0; i-- {
		wordAt := followingWordOffsets[i].word

		if len(strSoFar) > 0 {
			wordAt += " " + strSoFar
		}

		acc = traverseBreaks(followingWordOffsets[i].next, sChars, breaks, acc, wordAt)
	}

	return acc
}

type wordBreakNode struct {
	at   int
	word string
	next int
}

func wordBreak(s string, wordDict []string) []string {
	wordDictMap := make(map[string]bool)
	for _, word := range wordDict {
		wordDictMap[word] = true
	}

	sChars := []rune(s)
	n := len(sChars)
	// wordkBreakExists tracks whether the substring s[0:x] can be broken into valid words or not
	wordBreakExists := make([]bool, n+1)

	wordBreakExists[0] = true

	breaks := make([][]wordBreakNode, n+1)

	for i := 0; i <= n; i++ {
		breaks[i] = []wordBreakNode{}
	}

	for l := 1; l <= n; l++ {
		for j := 0; j < l; j++ {
			remainingWord := string(sChars[j:l])

			if wordBreakExists[j] && wordDictMap[remainingWord] {
				wordBreakExists[l] = true

				// Store the length of remainingWord so that we can print it later
				breaks[l] = append(breaks[l], wordBreakNode{at: l, word: remainingWord, next: j})
			}
		}
	}

	if wordBreakExists[n] {
		return traverseBreaks(n, sChars, breaks, []string{}, "")
	}

	return []string{}
}

type testCase struct {
	s              string
	dict           []string
	expectedResult []string
}

func tests() {
	testCases := []testCase{
		// testCase{"leetcode", []string{"leet", "code"}, []string{}},
		// testCase{"leetsscode", []string{"leet", "leets", "sscode"}, []string{}},
		// testCase{"applepenapple", []string{"apple", "pen"}, []string{}},
		// testCase{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, []string{}},
		testCase{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, []string{}},
		testCase{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}, []string{}},
	}

	for _, tc := range testCases {
		result := wordBreak(tc.s, tc.dict)

		fmt.Printf("result: %v\n", result)
	}
}

func main() {
	tests()
}
