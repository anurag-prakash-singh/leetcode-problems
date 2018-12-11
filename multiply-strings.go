package main

import (
	"bytes"
	"fmt"
)

// multiplicand should only comprise of digits
func multiplyByDigit(multiplicand []int, multiplier int, acc []int, accOffset int) {
	mulCarry := 0
	accCarry := 0

	for i, mulDigit := range multiplicand {
		prod := mulDigit*multiplier + mulCarry
		mulCarry = prod / 10
		prod %= 10
		acc[accOffset+i] += prod + accCarry
		accCarry = acc[accOffset+i] / 10
		acc[accOffset+i] %= 10
		// fmt.Printf("prod: %d; mulCarry = %d; accCarry: %d\n", prod, mulCarry, accCarry)
	}

	accCarry += mulCarry

	for i := (accOffset + len(multiplicand)); accCarry != 0; i++ {
		acc[i] += accCarry
		accCarry = acc[i] / 10
		acc[i] %= 10
	}
}

func multiply(num1 string, num2 string) string {
	num1Digits := make([]int, len(num1))
	num2Digits := make([]int, len(num2))
	digitRunes := map[int]rune{
		0: '0',
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
		9: '9',
	}

	for i, n1 := range num1 {
		num1Digits[len(num1)-1-i] = int(byte(n1) - '0')
	}

	for i, n2 := range num2 {
		num2Digits[len(num2)-1-i] = int(byte(n2) - '0')
	}

	acc := make([]int, len(num1)+len(num2)+5, len(num1)+len(num2)+5)

	for i, multiplierDigit := range num2Digits {
		multiplyByDigit(num1Digits, multiplierDigit, acc, i)
	}

	prodBuf := bytes.NewBufferString("")

	for i, skip := len(acc)-1, true; i >= 0; i-- {
		if acc[i] == 0 && skip {
			continue
		}

		skip = false

		prodBuf.WriteRune(digitRunes[acc[i]])
	}

	// fmt.Printf("product: %v\n", acc)

	ans := prodBuf.String()

	if ans == "" {
		return "0"
	}

	return prodBuf.String()
}

func main() {
	fmt.Printf("17 * 35: %s\n", multiply("17", "35"))
	fmt.Printf("172 * 543: %s\n", multiply("172", "543"))
	fmt.Printf("99999 * 99999: %s\n", multiply("99999", "99999"))
	fmt.Printf("1 * 99999: %s\n", multiply("1", "99999"))
	fmt.Printf("99999 * 1: %s\n", multiply("99999", "1"))
	fmt.Printf("0 * 0: %s\n", multiply("0", "0"))
}
