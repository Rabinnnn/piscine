package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	printDigits(n)
}

func printDigits(n int) {
	if n >= 10 {
		printDigits(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))
}
