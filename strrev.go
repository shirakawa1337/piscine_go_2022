package piscine

func StrRev(s string) string {
	r1 := []rune(s)
	reversedString := ""
	for i := len(r1) - 1; i >= 0; i-- {
		reversedString += string(r1[i])
	}
	return reversedString
}
