package main

import "fmt"

func longestPalindrome(s string) string {
	sChars := []rune(s)
	n := len(s)
	dp := make([][]bool, n+1)

	if n == 0 {
		return ""
	}

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n)
	}

	// All substrings of length 1 are palindromic by definition
	maxPalindromeStartOffset := 0
	maxPalindromeLength := 0

	for i := 0; i < n; i++ {
		dp[1][i] = true
		maxPalindromeLength = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			enclStart := i + 1
			enclEnd := i + l - 2

			if enclEnd >= n || enclStart >= n || i+l-1 >= n {
				break
			}

			if enclStart >= enclEnd {
				dp[l][i] = sChars[i] == sChars[i+l-1]
			} else {
				dp[l][i] = dp[l-2][i+1] && (sChars[i] == sChars[i+l-1])
			}

			if dp[l][i] {
				maxPalindromeStartOffset = i
				maxPalindromeLength = l
			}
		}
	}

	return string(sChars[maxPalindromeStartOffset : maxPalindromeStartOffset+maxPalindromeLength])
}

func tests() {
	testCases := []string{
		"babad",
		"cbbd",
		"",
		"a",
		"aaaaa",
		"",
	}

	for i := 0; i < len(testCases); i++ {
		fmt.Printf("tc #%d: result: %s\n", i, longestPalindrome(testCases[i]))
	}
}

func main() {
	tests()
}
