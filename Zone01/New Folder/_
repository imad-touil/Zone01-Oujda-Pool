package pescine

func SplitWhiteSpaces(s string) []string {
	var newS []string
	thisWord := ""
	for _, char := range s {
		if char == ' ' {
			if thisWord != "" {
				newS = append(newS, thisWord)
				thisWord = ""
			}
		} else {
			thisWord += string(char)
		}
	}
	if thisWord != "" {
		newS = append(newS, thisWord)
	}
	return newS
}
