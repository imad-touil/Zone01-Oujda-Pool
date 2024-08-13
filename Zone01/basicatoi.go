package pescine

func BasicAtoi(s string) int {
	result := 0
	for _, digit := range s {
		if digit >= '0' && digit <= '9' {
			result = result*10 + int(digit-'0')
		} else {
			return 0
		}
	}
	return result
}
