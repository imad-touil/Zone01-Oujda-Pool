package pescine

func Compact(ptr *[]string) int {
	com_arry := []string{}
	count := 0
	for _, char := range *ptr {
		if char != "" {
			com_arry = append(com_arry, char)
			count++
		}
	}
	*ptr = com_arry
	return count
}
