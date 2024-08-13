package pescine

func ToUpper(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] >= 'a' && r[i] <= 'z' {
			r[i] = r[i] - 32
		}
	}
	return string(r)
}
