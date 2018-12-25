package main

import "fmt"

var initedLetterOccurrenceTable map[rune]int
var initedLetterOccurrenceArr []int

func letterCount(str string) map[rune]int {
	strRunes := []rune(str)
	counts := initedLetterOccurrenceTable // newLetterOccurrenceTable()

	for _, r := range strRunes {
		counts[r]++
	}

	return counts
}

func letterCountArr(str string) []int {
	strRunes := []rune(str)
	counts := newLetterOccurrenceArr()

	for _, r := range strRunes {
		counts[r-'a']++
	}

	return counts
}

func initLetterOccurrenceTable() {
	initedLetterOccurrenceTable = make(map[rune]int)

	for c := 'a'; c <= 'z'; c++ {
		initedLetterOccurrenceTable[c] = 0
	}
}

func initLetterOccurrenceArr() {
	initedLetterOccurrenceArr = make([]int, 26, 26)

	for c := 'a'; c <= 'z'; c++ {
		initedLetterOccurrenceArr[c-'a'] = 0
	}
}

func clearInitLetterOccurrenceTable() {
	for c := 'a'; c <= 'z'; c++ {
		initedLetterOccurrenceTable[c] = 0
	}
}

func clearInitLetterOccurrenceArr() {
	for c := 'a'; c <= 'z'; c++ {
		initedLetterOccurrenceArr[c-'a'] = 0
	}
}

func newLetterOccurrenceTable() map[rune]int {
	letterMap := make(map[rune]int)

	for c := 'a'; c <= 'z'; c++ {
		letterMap[c] = 0
	}

	return letterMap
}

func newLetterOccurrenceArr() []int {
	letters := make([]int, 26, 26)

	for c := 'a'; c <= 'z'; c++ {
		letters[c-'a'] = 0
	}

	return letters
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func wordSubsets(A []string, B []string) []string {
	initLetterOccurrenceArr()
	maxBLetterCount := newLetterOccurrenceArr()

	for _, b := range B {
		clearInitLetterOccurrenceArr()
		lc := letterCountArr(b)

		for l, c := range lc {
			maxBLetterCount[l] = max(maxBLetterCount[l], c)
		}
	}

	universals := make([]string, 0, len(A))

	for _, a := range A {
		clearInitLetterOccurrenceArr()
		lc := letterCountArr(a)

		isUniversal := true

		for l, c := range lc {
			if c < maxBLetterCount[l] {
				isUniversal = false
				break
			}
		}

		if isUniversal {
			universals = append(universals, a)
		}
	}

	return universals
}

func test1() {
	a := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	b := []string{"e", "o"}

	fmt.Printf("wordSubsets: %v\n", wordSubsets(a, b))
}

func test2() {
	a := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	b := []string{"e", "oo"}

	fmt.Printf("wordSubsets: %v\n", wordSubsets(a, b))
}

func test3() {
	a := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	b := []string{"ec", "oc", "ceo"}

	fmt.Printf("wordSubsets: %v\n", wordSubsets(a, b))
}

func main() {

	// str := "abcdef"
	// runes := []rune(str)

	// fmt.Printf("letterMap: %v\n", newLetterOccurrenceTable())
	// fmt.Printf("runes: %v\n", runes)
	// fmt.Printf("counts: %v\n", letterCount("aabcaa"))

	test1()
	test2()
	test3()

}
