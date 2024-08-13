package pescine

func UltimateDivMod(a *int, b *int) {
	var div int = *a
	var mod int = *b
	if mod != 0 {
		*a = div / mod
		*b = div % mod
	}
}
