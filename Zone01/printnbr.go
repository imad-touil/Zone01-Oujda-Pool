package pescine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		PrintDigits(223372036854775808)
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	PrintDigits(n)
}

func PrintDigits(n int) {
	if n >= 10 {
		PrintDigits(n / 10)
	}

	z01.PrintRune(rune('0' + n%10))
}
