package piscine

func Split(s, sep string) []string {
	var result []string
	length := len(sep)
	current := 0

	for i := 0; i <= len(s)-length; i++ {
		if s[i:i+length] == sep {
			result = append(result, s[current:i])
			current = i + length
		}
	}

	result = append(result, s[current:])
	return result
}

