package main

import (
	"fmt"
)

func mismatchedParens(s string) (int, int) {
	lParens, rParens := 0, 0
	chars := []rune(s)
	mismatchedLParens, mismatchedRParens := 0, 0

	for _, c := range chars {
		if c == '(' {
			lParens++
		}

		if c == ')' {
			rParens++

			if lParens == 0 {
				mismatchedRParens++
			} else {
				lParens--
				rParens--
			}
		}

	}

	mismatchedLParens = lParens

	return mismatchedLParens, mismatchedRParens
}

func removeParensAt(chars []rune, at, lParensToDelete, rParensToDelete int, results map[string]struct{}) map[string]struct{} {
	if at >= len(chars) || at < 0 {
		if lParensToDelete == 0 && rParensToDelete == 0 {
			// Check if balanced and add to results if it is
			s := string(chars)
			mismatchedL, mismatchedR := mismatchedParens(s)

			if mismatchedL == 0 && mismatchedR == 0 {
				results[s] = struct{}{}
			}
		}

		return results
	}

	c := chars[at]

	if c != '(' && c != ')' {
		return removeParensAt(chars, at+1, lParensToDelete, rParensToDelete, results)
	}

	if c == '(' {
		if lParensToDelete > 0 {
			atCharDeleted := []rune{}
			atCharDeleted = append(atCharDeleted, chars[0:at]...)
			if at < len(chars)-1 {
				atCharDeleted = append(atCharDeleted, chars[at+1:]...)
			}

			results = removeParensAt(atCharDeleted, at, lParensToDelete-1, rParensToDelete, results)
		}
	} else {
		if rParensToDelete > 0 {
			atCharDeleted := []rune{}
			atCharDeleted = append(atCharDeleted, chars[0:at]...)
			if at < len(chars)-1 {
				atCharDeleted = append(atCharDeleted, chars[at+1:]...)
			}

			results = removeParensAt(atCharDeleted, at, lParensToDelete, rParensToDelete-1, results)
		}
	}

	unmodified := []rune(string(chars))
	results = removeParensAt(unmodified, at+1, lParensToDelete, rParensToDelete, results)

	return results
}

func removeInvalidParentheses(s string) []string {
	mismatchedL, mismatchedR := mismatchedParens(s)

	if mismatchedL == 0 && mismatchedR == 0 || len(s) == 0 {
		return []string{s}
	}

	results := make(map[string]struct{})
	results = removeParensAt([]rune(s), 0, mismatchedL, mismatchedR, results)

	resultList := []string{}
	for k, _ := range results {
		resultList = append(resultList, k)
	}

	return resultList
}

func testMismatchedParens() {
	testStrs := []string{
		"( ) ) ) ( ( ( )", // 0
		"(",               // 1
		")",               // 2
		"( ( )",           // 3
		"( ) )",           // 4
		"( ( ) )",         // 5
		"( ) ( )",         // 6
		"( ( ) ) )",       // 7
		"( ) ( ) ) ( )",   // 8
	}

	for i, ts := range testStrs {
		mL, mR := mismatchedParens(ts)

		fmt.Printf("Test %d: mL = %d, mR = %d\n", i, mL, mR)
	}
}

func testRemoveInvalidParens() {
	testStrs := []string{
		// "( ) ) ) ( ( ( )", // 0
		// "(",               // 1
		// ")",               // 2
		"(a(b)", // 3
		// "( ) )",           // 4
		"(())", // 5
		// "( ) ( )",         // 6
		// "( ( ) ) )",       // 7
		"(a)())()", // 8
		"()())()",  // 9
		")(",       // 10
	}

	for i, ts := range testStrs {
		results := removeInvalidParentheses(ts)

		fmt.Printf("Test %d:", i)

		for _, result := range results {
			fmt.Printf(" \"%s\" ", result)

		}

		fmt.Println()
	}
}

func main() {
	// testMismatchedParens()
	testRemoveInvalidParens()
}
