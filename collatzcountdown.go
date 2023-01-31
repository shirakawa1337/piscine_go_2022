package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	counter := 0

	for start > 0 {
		if start%2 == 0 {
			start /= 2
			counter++

		} else if start == 1 {
			break
		} else {
			start = (start * 3) + 1
			counter++
		}
	}
	return counter
}
