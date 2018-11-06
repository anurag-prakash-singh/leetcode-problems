package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	var result strings.Builder

	for num > 0 {
		if num <= 3 {
			result.WriteString("I")

			num--
		} else if num == 4 {
			result.WriteString("IV")

			num -= 4
		} else if num >= 5 && num <= 8 {
			result.WriteString("V")

			num -= 5
		} else if num == 9 {
			result.WriteString("IX")

			num -= 9
		} else if num >= 10 && num <= 39 {
			result.WriteString("X")

			num -= 10
		} else if num >= 40 && num <= 49 {
			result.WriteString("XL")

			num -= 40
		} else if num >= 50 && num <= 89 {
			result.WriteString("L")

			num -= 50
		} else if num >= 90 && num <= 99 {
			result.WriteString("XC")

			num -= 90
		} else if num >= 100 && num <= 399 {
			result.WriteString("C")

			num -= 100
		} else if num >= 400 && num <= 499 {
			result.WriteString("CD")

			num -= 400
		} else if num >= 500 && num <= 899 {
			result.WriteString("D")

			num -= 500
		} else if num >= 900 && num <= 999 {
			result.WriteString("CM")

			num -= 900
		} else if num >= 1000 {
			result.WriteString("M")

			num -= 1000
		}
	}

	return result.String()
}

func main() {

	tests := map[int]string{
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		6:    "VI",
		7:    "VII",
		9:    "IX",
		10:   "X",
		11:   "XI",
		14:   "XIV",
		15:   "XV",
		17:   "XVII",
		20:   "XX",
		23:   "XXIII",
		25:   "XXV",
		29:   "XXIX",
		30:   "XXX",
		39:   "XXXIX",
		40:   "XL",
		42:   "XLII",
		44:   "XLIV",
		49:   "XLIX",
		58:   "LVIII",
		85:   "LXXXV",
		90:   "XC",
		95:   "XCV",
		99:   "XCIX",
		1729: "MDCCXXIX",
		1994: "MCMXCIV",
		2434: "MMCDXXXIV",
		3442: "MMMCDXLII",
		3999: "MMMCMXCIX",
	}

	for n, s := range tests {
		ans := intToRoman(n)
		if ans != s {
			fmt.Printf("Test failed for %d. Expected %s; got %s\n", n, s, ans)
		}
	}

}
