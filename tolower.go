package piscine

func ToLower(s string) string {
	r1 := []rune(s)
	for i := 0; i < len(r1); i++ {
		if r1[i] >= 'a' && r1[i] <= 'z' {
			continue
		} else if r1[i] >= 'A' && r1[i] <= 'Z' {
			r1[i] = r1[i] + 32
		} else {
			continue
		}
	}
	return string(r1)
}
