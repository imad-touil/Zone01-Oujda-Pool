package pescine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var k []int
	if n <= 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		k = append(k, n%10)
		n = n / 10
	}
	for i := 0; i < len(k)-1; i++ {
		for j := i + 1; j < len(k); j++ {
			if k[j] < k[i] {
				k[j], k[i] = k[i], k[j]
			}
		}
	}
	for i := 0; i < len(k); i++ {
		z01.PrintRune(rune(k[i] + 48))
	}
}
