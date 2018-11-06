package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	sgn := 1
	num := 0
	numBegun := false
	numDigs := 0

	for i := 0; i < len(str); i++ {
		if numBegun && (str[i] < '0' || str[i] > '9') {
			break
		}

		if str[i] == ' ' {
			continue
		}

		if str[i] == '-' {
			sgn = -1
			numBegun = true

			continue
		}

		if str[i] == '+' {
			numBegun = true

			continue
		}

		if str[i] < '0' || str[i] > '9' {
			break
		}

		if num == -1*(1<<31) || num == ((1<<31)-1) {
			break
		}

		if str[i] >= '0' && str[i] <= '9' {
			numBegun = true
			digit := (int)(str[i] - '0')

			if numDigs == 10 {
				if sgn == -1 {
					// Hack ... we don't want -1 being multiplied with INT_MIN (the result will overflow INT_MAX)
					sgn = 1
					num = -2147483648
				} else {
					num = 2147483647
				}

				break
			}

			num = 10*num + digit
			numDigs = len(fmt.Sprintf("%d", num))
		}
	}

	ans := sgn * num

	if ans < math.MinInt32 {
		return math.MinInt32
	} else if ans > math.MaxInt32 {
		return math.MaxInt32
	} else {
		return ans
	}
}

func main() {

	tests := map[string]int{
		"1234":           1234,
		"  1234":         1234,
		"  ":             0,
		"  +":            0,
		"  -":            0,
		" 1 2":           1,
		" 1+2":           1,
		"-1+2":           -1,
		"++1":            0,
		"--1":            0,
		"abcd1234":       0,
		"  1234 43":      1234,
		"-91283472332":   -2147483648,
		"91283472332":    2147483647,
		" -91283472332 ": -2147483648,
		" 91283472332 ":  2147483647,
		"2147483648":     2147483647,
		"-2147483649":    -2147483648,
	}

	failures := make(map[string]int)
	failureMessages := make([]string, 0, len(tests))

	for str, num := range tests {
		ans := myAtoi(str)

		if ans != num {
			failureMessages = append(failureMessages, fmt.Sprintf("Test failed for %s. Expected %d, got %d\n", str, num, ans))
			failures[str] = ans
		}
	}

	if len(failures) == 0 {
		fmt.Printf("All tests passed\n")
	} else {
		for _, msg := range failureMessages {
			fmt.Printf("Failure: %s\n", msg)
		}
	}

}
