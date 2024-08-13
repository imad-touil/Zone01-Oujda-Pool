package pescine

func ToLower(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] >= 'A' && r[i] <= 'Z' {
			r[i] = r[i] + 32
		}
	}
	return string(r)
}
