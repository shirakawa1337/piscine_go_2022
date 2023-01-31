package piscine

func AlphaCount(s string) int {
	r1 := []rune(s)
	nb := 0
	for i := 0; i < len(r1); i++ {
		if (rune('a') <= r1[i] && r1[i] <= rune('z')) ||
			(rune('A') <= r1[i] && r1[i] <= rune('Z')) {
			nb++
		}
	}
	return nb
}
