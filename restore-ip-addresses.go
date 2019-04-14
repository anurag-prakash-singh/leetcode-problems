package main

import (
	"fmt"
	"strconv"
)

func possibleIPs(rawIP []rune, start int, dotOffsets []int, results []string) []string {
	if len(dotOffsets) > 3 {
		return results
	}

	if start >= len(rawIP) {
		return results
	}

	if len(dotOffsets) == 3 && dotOffsets[len(dotOffsets)-1] < len(rawIP)-1 {
		// Form string, add to results
		newIP := []rune{}

		for i, dotOffset := range dotOffsets {
			if i == 0 {
				newIP = append(newIP, rawIP[0:dotOffset+1]...)
			} else {
				newIP = append(newIP, rawIP[dotOffsets[i-1]+1:dotOffset+1]...)
			}

			if i < 3 {
				newIP = append(newIP, '.')
			}
		}

		lastDotOffset := dotOffsets[2]
		remChars := rawIP[lastDotOffset+1:]

		if len(remChars) > 3 || len(remChars) == 0 || (len(remChars) > 1 && remChars[0] == '0') {
			return results
		}

		if remNum, err := strconv.Atoi(string(remChars)); err != nil {
			return results
		} else if remNum > 255 {
			return results
		}

		newIP = append(newIP, remChars...)

		results = append(results, string(newIP))

		return results
	}

	digit1 := rawIP[start] - '0'

	dotOffsets = append(dotOffsets, start)
	results = possibleIPs(rawIP, start+1, dotOffsets, results)
	dotOffsets = dotOffsets[0 : len(dotOffsets)-1]

	if start+1 >= len(rawIP) {
		return results
	}

	digit2 := rawIP[start+1] - '0'

	if digit1 == 0 {
		return results
	}

	dotOffsets = append(dotOffsets, start+1)
	results = possibleIPs(rawIP, start+2, dotOffsets, results)
	dotOffsets = dotOffsets[0 : len(dotOffsets)-1]

	if start+2 >= len(rawIP) {
		return results
	}

	digit3 := rawIP[start+2] - '0'

	threeDigNum := 100*digit1 + 10*digit2 + digit3

	if threeDigNum > 255 {
		return results
	}

	dotOffsets = append(dotOffsets, start+2)
	results = possibleIPs(rawIP, start+3, dotOffsets, results)
	dotOffsets = dotOffsets[0 : len(dotOffsets)-1]

	return results
}

func restoreIpAddresses(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	rawIP := []rune(s)

	return possibleIPs(rawIP, 0, []int{}, []string{})
}

func test1() {
	ans := restoreIpAddresses("1234")

	fmt.Printf("ans = %v\n", ans)
}

func test2() {
	ans := restoreIpAddresses("25525511135")

	fmt.Printf("ans = %v\n", ans)
}

func tests() {
	rawIPs := []string{
		"1234",
		"25525511135",
		"1238",
		"8888",
		"0000",
		"000",
		"",
		"0",
		"010010",
	}

	for _, rawIP := range rawIPs {
		ans := restoreIpAddresses(rawIP)

		fmt.Printf("rawIP = %s; ans = %v\n", rawIP, ans)
	}
}

func main() {
	// test1()
	// test2()
	tests()
}
