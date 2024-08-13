package pescine

func IsNumeric(s string) bool {
	for digit := 0; digit < len(s); digit++ {
		if s[digit] < '0' || s[digit] > '9' {
			return false
		}
	}
	return true
}
