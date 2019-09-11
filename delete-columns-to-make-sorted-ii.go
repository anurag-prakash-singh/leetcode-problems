package main

import "fmt"

func minDeletionSize(A []string) int {
	charss := make([][]rune, len(A))

	for i, _ := range A {
		charss[i] = []rune(A[i])
	}

	strictlyGreater := make([]bool, len(A))
	numDeletions := 0

	for j := 0; j < len(A[0]); j++ {
		monotonicIncreasing := true

		for i := 1; i < len(A); i++ {
			if charss[i-1][j] >= charss[i][j] {
				monotonicIncreasing = false
			}

			if charss[i-1][j] > charss[i][j] {
				if strictlyGreater[i] {
					continue
				}

				// inversion
				if j == 0 {
					// inversion in first column
					numDeletions++

					for k := 0; k < len(A); k++ {
						charss[k][j] = '_'
					}

					break
				} else {
					// inversion that's not in first column
					if charss[i-1][j-1] == charss[i][j-1] {
						// copy previous column and delete this one
						numDeletions++

						for k := 0; k < len(A); k++ {
							charss[k][j] = charss[k][j-1]
						}

						break
					}
				}
			}
		}

		for i := 1; i < len(A); i++ {
			if charss[i-1][j] < charss[i][j] {
				strictlyGreater[i] = true
			}
		}

		if monotonicIncreasing {
			return numDeletions
		}
	}

	return numDeletions
}

func test1() {
	A := []string{
		"ca",
		"bb",
		"ac",
	}

	fmt.Printf("minDeletions: %d\n", minDeletionSize(A))
}

func test2() {
	A := []string{
		"xc",
		"yb",
		"za",
	}

	fmt.Printf("minDeletions: %d\n", minDeletionSize(A))
}

func test3() {
	A := []string{
		"zyx",
		"wvu",
		"tsr",
	}

	fmt.Printf("minDeletions: %d\n", minDeletionSize(A))
}

func main() {
	// test1()
	// test2()
	test3()
}
