package pescine

func StrRev(s string) string {
	var newS string
	for i := len(s) - 1; i >= 0; i-- {
		newS += string(s[i])
	}
	return newS
}
