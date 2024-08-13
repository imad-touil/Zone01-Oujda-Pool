package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	upperCase := false

	if len(os.Args) > 1 && os.Args[1] == "--upper" {
		upperCase = true
		os.Args = os.Args[1:]
	}

	for _, arg := range os.Args[1:] {
		n := 0
		for i := 0; i < len(arg); i++ {
			if arg[i] < '0' || arg[i] > '9' {
				n = -1
				break
			}
			n = n*10 + int(arg[i]-'0')
		}

		if n < 1 || n > 26 {
			z01.PrintRune(' ')
		} else {
			letter := 'a' + n - 1
			if upperCase {
				letter = 'A' + n - 1
			}
			z01.PrintRune(rune(letter))
		}
	}
	z01.PrintRune('\n')
}
