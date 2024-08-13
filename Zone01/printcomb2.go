package pescine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for m := '0'; m <= '9'; m++ {
			d := m + 1
			for a := i; a <= '9'; a++ {
				for ; d <= '9'; d++ {
					z01.PrintRune(i)
					z01.PrintRune(m)
					z01.PrintRune(' ')
					z01.PrintRune(a)
					z01.PrintRune(d)
					if i < '9' || m < '8' || a < '9' || d < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				d = '0'
			}
		}
	}
	z01.PrintRune('\n')
}
