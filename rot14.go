package piscine

func Rot14(s string) string {
	r1 := []rune(s)
	for i := 0; i < len(r1); i++ {
		if r1[i] >= 'a' && r1[i] <= 'z' {
			if r1[i] >= 'm' {
				r1[i] -= 12
			} else {
				r1[i] += 14
			}
		} else if r1[i] >= 'A' && r1[i] <= 'Z' {
			if r1[i] >= 'M' {
				r1[i] -= 12
			} else {
				r1[i] += 14
			}
		} else {
			continue
		}
	}
	return string(r1)
}
