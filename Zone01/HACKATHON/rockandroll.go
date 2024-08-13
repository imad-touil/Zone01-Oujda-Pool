package pescine

func RockAndRoll(n int) string {
	rock := "rock\n"
	roll := "roll\n"
	rr := "rock and roll\n"
	err1 := "error: non divisible\n"
	err2 := "error: number is negative\n"
	if n < 0 {
		return err2
	}
	if n%2 == 0 && n%3 == 0 {
		return rr
	}
	if n%2 == 0 {
		return rock
	}
	if n%3 == 0 {
		return roll
	}
	return err1
}
