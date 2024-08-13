package pescine

func IsAlpha(s string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if (s[i] < '0' || s[i] > '9') && (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') {
			return false
		}
	}
	return true
}
