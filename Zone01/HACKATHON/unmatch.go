package pescine

func Unmatch(a []int) int {
	var counter int
	for _, val1 := range a {
		counter = 0
		for _, val2 := range a {
			if val2 == val1 {
				counter++
			}
		}
		if counter%2 != 0 {
			return val1
		}
	}
	return -1
}
