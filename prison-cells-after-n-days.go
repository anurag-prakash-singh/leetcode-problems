package main

import "fmt"

func setBit(num, position uint8) uint8 {
	return (num | (1 << position))
}

func resetBit(num, position uint8) uint8 {
	return (num &^ (1 << position))
}

func getBit(num, position uint8) uint8 {
	if (num & (1 << position)) == 0 {
		return 0
	}

	return 1
}

func calculateNext(num uint8) uint8 {
	nextNum := uint8(0)

	for i := uint8(1); i < 7; i++ {
		if getBit(num, i-1) == 0 && getBit(num, i+1) == 0 {
			nextNum = setBit(nextNum, i)

			continue
		}

		if getBit(num, i-1) == 1 && getBit(num, i+1) == 1 {
			nextNum = setBit(nextNum, i)

			continue
		}

		nextNum = resetBit(nextNum, i)
	}

	return nextNum
}

func slcToNum(bits []int) uint8 {
	var num uint8

	for _, b := range bits {
		num = num << 1
		num |= uint8(b)
	}

	return num
}

func numToSlc(num uint8) []int {
	slc := make([]int, 8)

	for i := uint8(0); i < 8; i++ {
		slc[i] = int(getBit(num, 7-i))
	}

	return slc
}

func prisonAfterNDays(cells []int, N int) []int {
	num := slcToNum(cells)
	lastSeen := make([]int, 256)

	for i := 0; i < len(lastSeen); i++ {
		lastSeen[i] = -1
	}

	lastSeen[num] = 0

	checkedOffset := 1

	for checkedOffset <= N {
		num = calculateNext(num)

		if lastSeen[num] != -1 || checkedOffset == N {
			// Collision detected
			fmt.Printf("Collition detected on day %d: lastSeen[num]: %d\n", checkedOffset, lastSeen[num])
			break
		}

		lastSeen[num] = checkedOffset

		checkedOffset++
	}

	if checkedOffset == N {
		return numToSlc(num)
	}

	if lastSeen[num] == 0 {
		N = N % (checkedOffset - lastSeen[num])
	} else {
		N = (N % (checkedOffset - lastSeen[num])) + (checkedOffset - lastSeen[num] - 1)
	}

	fmt.Printf("N: %d; num: %v\n", N, numToSlc(num))

	for i := 1; i <= N; i++ {
		num = calculateNext(num)
	}

	return numToSlc(num)
}

func test1() {
	cells := []int{0, 1, 0, 1, 1, 0, 0, 1}
	// num := slcToNum(cells)
	// fmt.Printf("slc: %v\n", numToSlc(num))

	// printNum(num)

	// nextNum := calculateNext(num)
	// printNum(nextNum)
	// nextNum = calculateNext(nextNum)
	// printNum(nextNum)
	fmt.Printf("ans: %v\n", prisonAfterNDays(cells, 7))

}

func test2() {
	cells := []int{1, 0, 0, 1, 0, 0, 1, 0}

	fmt.Printf("ans: %v\n", prisonAfterNDays(cells, 1000000000))
}

func test3() {
	cells := []int{0, 0, 1, 1, 1, 1, 0, 0}

	fmt.Printf("ans: %v\n", prisonAfterNDays(cells, 8))
}

func test4() {
	cells := []int{1, 0, 0, 1, 0, 0, 0, 1}

	fmt.Printf("ans: %v\n", prisonAfterNDays(cells, 826))
}

func printNum(num uint8) {
	for i := uint8(7); true; i-- {
		fmt.Printf("%d ", getBit(num, i))

		if i == 0 {
			break
		}
	}

	fmt.Println()
}

func main() {
	// fmt.Printf("%d\n", setBit(0, 5))
	// test1()
	test2()
	// test3()
	// test4()
}
