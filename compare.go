package piscine

func Compare(a, b string) int {
	r1 := []rune(a)
	r2 := []rune(b)

	minLen := len(r1)
	if minLen > len(r2) {
		minLen = len(r2)
	}

	for i := 0; i < minLen; i++ {
		if r1[i] > r2[i] {
			return 1
		} else if r1[i] < r2[i] {
			return -1
		}
	}

	if len(r1) > len(r2) {
		return 1
	} else if len(r1) < len(r2) {
		return -1
	} else {
		return 0
	}
}
