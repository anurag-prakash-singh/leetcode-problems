package main

import (
	"fmt"
)

var pad = map[int]string{
	0: " ",
	1: "",
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func enumerateCombins(digits string, digitsOffset int, currResult string, resultsSoFar []string) []string {
	if digitsOffset == len(digits) {
		return append(resultsSoFar, currResult)
	}

	digitNum := (int)(digits[digitsOffset] - '0')

	for _, digitLetter := range pad[digitNum] {
		resultsSoFar = enumerateCombins(digits, digitsOffset+1, currResult+string(digitLetter), resultsSoFar)
	}

	return resultsSoFar
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	result := make([]string, 0, 10)

	result = enumerateCombins(digits, 0, "", result)

	return result
}

func main() {
	result1 := letterCombinations("")

	fmt.Printf("result: %v\n", result1)
}
