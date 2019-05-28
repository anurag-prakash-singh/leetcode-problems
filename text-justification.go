package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	justified := []string{}

	for i := 0; i < len(words); {
		j := i
		unjustifiedLineWordLen := 0

		for totLen := 0; j < len(words); j++ {
			totLen += len(words[j])

			if j > i {
				totLen += 1 // add 1 to account for the preceding space
			}

			if totLen <= maxWidth {
				unjustifiedLineWordLen += len(words[j])
			} else {
				// fmt.Printf("totLen: %d; j: %d\n", totLen, j)
				break
			}
		}

		// fmt.Printf("Ending word offset (incl): %d\n", j-1)

		// All words from [i] to [j - 1] (inclusive) should fit in 1 line
		slots := j - i - 1

		if slots <= 0 {
			// We can fit just 1 word
			justified = append(justified, words[i]+strings.Repeat(" ", maxWidth-len(words[i])))
		} else {
			perSlotSpacesFloor := (maxWidth - unjustifiedLineWordLen) / slots
			leftoverSpaces := (maxWidth - unjustifiedLineWordLen) % slots
			newLine := ""

			for wordOffset := i; wordOffset < j; wordOffset++ {
				newLine += words[wordOffset]

				if wordOffset < j-1 {
					if j >= len(words) {
						newLine += " "

						continue
					}

					newLine += strings.Repeat(" ", perSlotSpacesFloor)

					if leftoverSpaces > 0 {
						leftoverSpaces--
						newLine += " "
					}
				}
			}

			if len(newLine) < maxWidth {
				newLine += strings.Repeat(" ", maxWidth-len(newLine))
			}

			justified = append(justified, newLine)
		}

		i = j
	}

	return justified
}

func printStrings(strs []string) {
	for _, str := range strs {
		fmt.Printf("\"%s\"\n", str)
	}
}

type testCase struct {
	words    []string
	maxWidth int
}

func tests() {
	testCases := []testCase{
		testCase{words: []string{"This", "is", "an", "example", "of", "text", "justification."}, maxWidth: 16},
		testCase{words: []string{"What", "must", "be", "acknowledgment", "shall", "be"}, maxWidth: 16},
		testCase{words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, maxWidth: 20},
	}

	for i, tc := range testCases {
		justified := fullJustify(tc.words, tc.maxWidth)

		fmt.Printf("Output of test %d\n", i)
		printStrings(justified)
		fmt.Printf("---\n")
	}
}

func main() {
	tests()
}
