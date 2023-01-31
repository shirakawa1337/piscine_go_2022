package piscine

func isAlphaNumerical(value rune) bool {
	if (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') || (value >= '0' && value <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	arr := []rune(s)
	first := true

	for i := range arr {
		if isAlphaNumerical(arr[i]) == true && first {
			if arr[i] >= 'a' && arr[i] <= 'z' {
				arr[i] -= 32
			}
			first = false
		} else if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] += 32
		} else if isAlphaNumerical(arr[i]) == false {
			first = true
		}
	}
	return string(arr)
}
