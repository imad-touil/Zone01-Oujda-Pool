package pescine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	s := strs[0]
	for i := 1; i < len(strs); i++ {
		s += sep + strs[i]
	}
	return s
}
