package piscine

func IsAlpha(s string) bool {
	r1 := []rune(s)
	for i := 0; i < len(r1); i++ {
		if rune('a') <= r1[i] && rune('z') >= r1[i] {
			continue
		} else if rune('A') <= r1[i] && rune('Z') >= r1[i] {
			continue
		} else if rune('0') <= r1[i] && rune('9') >= r1[i] {
			continue
		} else {
			return false
		}
	}
	return true
}
