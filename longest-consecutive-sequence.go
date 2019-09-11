package main

import "fmt"

func max2(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numIndices := make(map[int]int)

	for i, n := range nums {
		numIndices[n] = i
	}

	distinctNums := make([]int, 0, len(nums))

	for k, _ := range numIndices {
		distinctNums = append(distinctNums, k)
	}

	for i, n := range distinctNums {
		numIndices[n] = i
	}

	consecStartLens := make(map[int]int)
	path := make([]int, 0, len(distinctNums))

	for k, _ := range numIndices {
		_, ok := numIndices[k+1]
		if !ok {
			consecStartLens[k] = 1
		}
	}

	maxConsecLen := 1

	for k, _ := range numIndices {
		curr := k
		kPath := path[:]

		for {
			consecLen := consecStartLens[curr]

			if consecLen >= 1 {
				// We've reached a point starting from which we know how many
				// consecutive elements follow
				kPathLen := len(kPath)
				for i, kp := range kPath {
					consecStartLens[kp] = kPathLen - i + consecLen
					maxConsecLen = max2(maxConsecLen, consecStartLens[kp])
				}

				break
			} else {
				kPath = append(kPath, curr)
			}

			curr += 1
		}
	}

	return maxConsecLen
}

func test1() {
	nums := []int{100, 4, 200, 1, 3, 2}
	result := longestConsecutive(nums)

	fmt.Printf("result: %d\n", result)
}

func test2() {
	nums := []int{100, -4, 200, -1, 3, 2}
	result := longestConsecutive(nums)

	fmt.Printf("result: %d\n", result)
}

func test3() {
	nums := []int{100, -4, 200}
	result := longestConsecutive(nums)

	fmt.Printf("result: %d\n", result)
}

func main() {
	// test1()
	// test2()
	test3()
}
