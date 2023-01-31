package piscine

func LastRune(s string) rune {
	r1 := []rune(s)
	lastRune := r1[len(r1)-1]
	return lastRune
}
