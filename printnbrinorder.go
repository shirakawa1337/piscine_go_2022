package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	var cnt [10]int
	for n > 0 {
		x := n % 10
		n /= 10
		cnt[x] += 1
	}
	for i := 0; i <= 9; i++ {
		for j := 1; j <= cnt[i]; j++ {
			z01.PrintRune(rune(i + '0'))
		}
	}
}
