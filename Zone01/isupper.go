package pescine

func IsUpper(s string) bool {
	for letter := 0; letter < len(s); letter++ {
		if s[letter] <= 'A' || s[letter] >= 'Z' {
			return false
		}
	}
	return true
}
