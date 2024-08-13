package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg)-1; i++ {
		for j := i + 1; j < len(arg); j++ {
			if arg[j] < arg[i] {
				arg[j], arg[i] = arg[i], arg[j]
			}
		}
	}
	for _, char := range arg {
		PrintStr(char)
		z01.PrintRune('\n')
	}
}

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}
