package main

import (
	"fmt"
)

func divide(dividend int, divisor int) int {

	if dividend == divisor {
		return 1
	}

	if divisor == 1 {
		return dividend
	}

	sgn := 1

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sgn = -1
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	if dividend < divisor {
		return 0
	}

	ans := 0
	temp := 0

	for i := uint32(31); ; i-- {
		if temp+(divisor<<i) <= dividend {
			temp += divisor << i
			ans = ans | (1 << i)
		}

		if i == 0 {
			break
		}
	}

	if sgn == -1 {
		ans = -ans
	}

	if ans >= 1<<31 {
		ans = (1 << 31) - 1
	}

	return ans
}

func main() {

	tests := [][3]int{{1, 2, 0}, {10, 3, 3}, {5, 3, 1}, {0, -5, 0}, {-2147483648, -1, 2147483647}, {-2147483647, 2, -2147483647 / 2}}

	for _, test := range tests {
		dividend, divisor, expAns := test[0], test[1], test[2]

		ans := divide(dividend, divisor)

		if ans != expAns {
			fmt.Printf("test failed: %d / %d should have been %d. Was %d.\n", dividend, divisor, expAns, ans)
		}
	}

}
