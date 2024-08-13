package pescine

func ConcatParams(args []string) string {
	s := args[0]
	for i := 1; i < len(args); i++ {
		s += "\n" + args[i]
	}
	return s
}
