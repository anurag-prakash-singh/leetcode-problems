package main

func min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func min3(a, b, c int) int {
	return min2(a, min2(b, c))
}

func minDistance(word1 string, word2 string) int {
	w1Chars, w2Chars := []rune(word1), []rune(word2)
	w1Len, w2Len := len(w1Chars), len(w2Chars)
	mat := make([][]int, w1Len+1)

	for i, _ := range mat {
		mat[i] = make([]int, w2Len+1)
	}

	// mat[i][j]: the minimum number of edits neeed to transform
	// word1's i-prefix (prefix of length i) to word2's j-prefix
	// (prefix of length j)
	for i := 0; i <= w1Len; i++ {
		mat[i][0] = i
	}

	for j := 0; j <= w2Len; j++ {
		mat[0][j] = j
	}

	for i := 1; i <= w1Len; i++ {
		for j := 1; j <= w2Len; j++ {
			if w1Chars[i-1] == w2Chars[j-1] {
				mat[i][j] = mat[i-1][j-1]

				continue
			}

			mat[i][j] = min3(1+mat[i-1][j-1], 1+mat[i-1][j], 1+mat[i][j-1])
		}
	}

	return mat[w1Len][w2Len]
}

func main() {

}
