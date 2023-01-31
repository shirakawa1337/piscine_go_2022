package piscine

func SplitWhiteSpaces(s string) []string {
	var output []string
	var tmp string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(tmp) > 0 {
				output = append(output, tmp)
			}
			tmp = ""
		} else {
			tmp += string(s[i])
		}
	}
	if len(tmp) > 0 {
		output = append(output, tmp)
	}
	return output
}
