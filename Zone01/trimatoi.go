package pescine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	i := 0
	for i < len(s) && (s[i] < '0' || s[i] > '9') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			for i < len(s) && (s[i] >= '0' && s[i] <= '9') {
				result = result*10 + int(s[i]-'0')
				i++
			}
		} else {
			i++
		}
	}
	return result * sign
}
