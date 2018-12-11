package main

import "fmt"

func newLetterOccurrenceTable() map[rune]int {
	letterMap := make(map[rune]int)

	for c := 'a'; c <= 'z'; c++ {
		letterMap[c] = 0
	}

	return letterMap
}

func main() {

	fmt.Printf("letterMap: %v\n", newLetterOccurrenceTable())

}
