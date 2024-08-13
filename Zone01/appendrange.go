package pescine

func AppendRange(min, max int) []int {
	size := 0
	var integer []int
	if max > min {
		size = max - min
		for i := 0; i < size; i++ {
			integer = append(integer, min)
			min++
		}
		return integer
	} else {
		return integer
	}
}
