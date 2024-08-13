package pescine

func ActiveBits(n int) int {
	counter := 0
	for n != 0 {
		counter += n % 2
		n = n / 2
	}
	return counter
}
