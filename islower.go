package piscine

func IsLower(s string) bool {
	r1 := []rune(s)
	for i := 0; i < len(r1); i++ {
		if rune('a') <= r1[i] && rune('z') >= r1[i] {
			continue
		} else {
			return false
		}
	}
	return true
}
