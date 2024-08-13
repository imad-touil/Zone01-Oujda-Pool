package pescine

func StringToIntSlice(str string) []int {
	var slicce []int

	for _, char := range str {
		slicce = append(slicce, int(char))
	}
	return slicce
}
