package main

import (
	"errors"
	"fmt"
)

func slicesMatch(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func runeExists(runes []rune, r rune) bool {
	for _, c := range runes {
		if c == r {
			return true
		}
	}

	return false
}

func topologicalOrder(starting rune, adjMat map[rune][]rune, visitedInPriorTraversal, visited map[rune]bool, orderSoFar string) (string, error) {
	if visited[starting] {
		return "", errors.New("already visited " + string(starting))
	}

	visited[starting] = true

	adjs := adjMat[starting]

	// if !ok || len(adjs) == 0 {
	// 	return orderSoFar + string(starting), nil
	// }

	for _, a := range adjs {
		if !visitedInPriorTraversal[a] {
			var err error
			orderSoFar, err = topologicalOrder(a, adjMat, visitedInPriorTraversal, visited, orderSoFar)

			if err != nil {
				return "", err
			}
		}
	}

	return orderSoFar + string(starting), nil
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func alienOrder(words []string) string {
	inEdges := make(map[rune]int)
	alphabet := make(map[rune]bool)
	wordChars := make([][]rune, len(words))
	maxWordLen := 0

	for i, str := range words {
		wordChars[i] = []rune(str)

		if len(str) > maxWordLen {
			maxWordLen = len(str)
		}

		for _, c := range wordChars[i] {
			inEdges[c] = 0
			alphabet[c] = true
		}
	}

	adjMat := make(map[rune][]rune)

	for letterOffset := 0; letterOffset < maxWordLen; letterOffset++ {
		for wordOffset := 0; wordOffset < len(words); wordOffset++ {
			if len(words[wordOffset]) <= letterOffset {
				continue
			}

			wordChar := wordChars[wordOffset][letterOffset]
			_, ok := adjMat[wordChar]

			if !ok {
				adjMat[wordChar] = []rune{}
			}

			if wordOffset == 0 {
				continue
			}

			if letterOffset > 0 && (letterOffset >= len(wordChars[wordOffset-1]) || !slicesMatch(wordChars[wordOffset][0:letterOffset], wordChars[wordOffset-1][0:letterOffset])) {
				continue
			}

			preChar := wordChars[wordOffset-1][letterOffset]

			if !runeExists(adjMat[preChar], wordChar) && preChar != wordChar {
				adjMat[preChar] = append(adjMat[preChar], wordChar)
				inEdges[wordChar]++
			}
		}
	}

	// for k, v := range adjMat {
	// 	fmt.Printf("%c -> %s\n", k, string(v))
	// }

	visitedInPriorTraversal := make(map[rune]bool)
	orderSoFar := ""

	for k, _ := range alphabet {
		if inEdges[k] > 0 {
			// We only want to start traversing from root characters
			continue
		}

		if !visitedInPriorTraversal[k] {
			var err error
			visited := map[rune]bool{}
			orderSoFar, err = topologicalOrder(k, adjMat, visitedInPriorTraversal, visited, orderSoFar)

			for visitedK, visitedV := range visited {
				visitedInPriorTraversal[visitedK] = visitedV
			}

			if orderSoFar == "" || err != nil {
				return ""
			}
		}
	}

	return reverse(orderSoFar)
}

func tests() {
	testCases := [][]string{
		[]string{"z", "x"},
		[]string{"wrt", "wrf", "er", "ett", "rftt"},
		[]string{"wrt", "wrfb", "wrff", "er", "ett", "rftt"},
		[]string{"z", "x", "z"},
		[]string{"abc"},
		[]string{"a"},
		[]string{""},
		[]string{"ab", "d", "b"},
	}

	for i, tc := range testCases {
		result := alienOrder(tc)

		fmt.Printf("testcase %d: result: %s\n", i, result)
	}
}

func main() {
	tests()
}
