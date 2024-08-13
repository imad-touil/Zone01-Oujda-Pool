package pescine

func ReverseMenuIndex(menu []string) []string {
	slice := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		slice[i] = menu[len(menu)-1-i]
	}
	return slice
}
