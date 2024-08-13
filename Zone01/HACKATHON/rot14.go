package main

import "github.com/01-edu/z01"

func Rot14(s string) string {
	str := ""
	for _, char := range s {
		if char >= 'A' && char < 'M' {
			str += string(char + 14)
		} else if char >= 'M' && char < 'A' {
			str += string(char - 12)
		} else if char == 'Z' {
			str += "N"
		} else if char >= 'a' && char < 'm' {
			str += string(char + 14)
		} else if char >= 'm' && char < 'z' {
			str+= string(char - 12)
		} else if char == 'z' {
			str += "n"
		} else {
			str += string(char)
		}
	}
	return str
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
