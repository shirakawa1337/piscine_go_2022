package piscine

func IsPrintable(s string) bool {
	r1 := []rune(s)
	for i := 0; i < len(r1); i++ {
		if r1[i] >= 0 && r1[i] <= 31 {
			return false
		} else {
			continue
		}
	}
	return true
}
