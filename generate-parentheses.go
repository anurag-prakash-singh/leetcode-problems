package main

import (
	"fmt"
)

func genParens(strSoFar string, remainingLParens, unpairedLParens int, results []string) []string {
	if remainingLParens == 0 {
		if len(strSoFar) == 0 {
			return results
		}

		if unpairedLParens > 0 {
			return genParens(strSoFar+string(")"), remainingLParens, unpairedLParens-1, results)
		}

		return append(results, strSoFar)
	}

	results = genParens(strSoFar+string("("), remainingLParens-1, unpairedLParens+1, results)

	if unpairedLParens > 0 {
		results = genParens(strSoFar+string(")"), remainingLParens, unpairedLParens-1, results)
	}

	return results
}

func generateParenthesis(n int) []string {
	results := make([]string, 0, 10)

	return append(results, genParens("", n, 0, results)...)
}

func main() {

	ans1 := generateParenthesis(3)

	fmt.Printf("ans: %v\n", ans1)

}
