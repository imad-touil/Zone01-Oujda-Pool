package pescine

func MakeRange(min, max int) []int {
	size := 0
	var integer []int
	if max > min {
		size = max - min
		integer := make([]int, size)
		for i := 0; i < size; i++ {
			integer[i] = min
			min++
		}
		return integer
	} else {
		return integer
	}
}
