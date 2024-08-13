package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	words := os.Args[1:]
	for i := len(words) - 1; i >= 0; i-- {
		PrintStr(words[i])
		z01.PrintRune('\n')
	}
}

func PrintStr(s string) {
	runes := []rune(s)
	for _, letter := range runes {
		z01.PrintRune(letter)
	}
}
