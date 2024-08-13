package pescine

func DescendAppendRange(max, min int) []int {
	val := []int{}
	if max <= min {
		return val
	} else {
		for i := max; i > min; i-- {
			val = append(val, i)
		}
		return val
	}
}
