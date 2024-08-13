package pescine

func Abort(a, b, c, d, e int) int {
	nbrs := [5]int{a, b, c, d, e}
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if nbrs[i] > nbrs[j] {
				nbrs[i], nbrs[j] = nbrs[j], nbrs[i]
			}
		}
	}
	return nbrs[2]
}
