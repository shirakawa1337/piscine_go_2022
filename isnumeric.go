package piscine

func IsNumeric(s string) bool {
	r1 := []rune(s)
	for i := 0; i < len(r1); i++ {
		if rune('0') <= r1[i] && rune('9') >= r1[i] {
			continue
		} else {
			return false
		}
	}
	return true
}
