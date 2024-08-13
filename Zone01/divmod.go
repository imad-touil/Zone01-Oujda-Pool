package pescine

func DivMod(a int, b int, div *int, mod *int) {
	if b != 0 {
		*div = a / b
		*mod = a % b
	}
}
