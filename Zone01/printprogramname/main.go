package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	for _, char := range name {
		if char == '.' || char == '/' {
			char++
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
