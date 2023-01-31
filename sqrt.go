package piscine

func Sqrt(nb int) int {
	output := 1
	output = nb / 2
	if output*output == nb {
		return output
	} else if nb == 1 {
		return 1
	} else {
		return 0
	}
}
