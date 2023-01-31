package piscine

func IsUpper(s string) bool {
	r1 := []rune(s)
	for i := 0; i < len(r1); i++ {
		if rune('A') <= r1[i] && rune('Z') >= r1[i] {
			continue
		} else {
			return false
		}
	}
	return true
}
