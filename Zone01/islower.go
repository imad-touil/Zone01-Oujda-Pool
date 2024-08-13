package pescine

func IsLower(s string) bool {
	for letter := 0; letter < len(s); letter++ {
		if s[letter] < 'a' || s[letter] > 'z' {
			return false
		}
	}
	return true
}
