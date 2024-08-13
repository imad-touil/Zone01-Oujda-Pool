package pescine

func IterativeFactorial(nb int) int {
	fac := 1
	if nb < 0 || nb > 100 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			fac = fac * i
		}
	}
	return fac
}
