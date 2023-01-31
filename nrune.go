package piscine

func NRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	}
	r1 := []rune(s)
	output := r1[n-1]
	return output
}
