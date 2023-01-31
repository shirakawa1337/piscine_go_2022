package piscine

func TrimAtoi(s string) int {
	r1 := []rune(s)
	nb := false
	sign := 0
	result := 0

	for i := 0; i < len(r1); i++ {
		if r1[i] >= '0' && r1[i] <= '9' {
			nb = true
			result *= 10
			result += int(r1[i] - '0')
		} else if r1[i] == '-' || r1[i] == '+' {
			if !nb && sign == 0 {
				if r1[i] == '-' {
					sign = -1
				} else if sign == 0 {
					sign = 1
				}
			}
		}
	}
	if sign == 0 {
		sign = 1
	}
	return result * sign
}
