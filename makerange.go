package piscine

func MakeRange(min, max int) []int {
	if min == max || min > max {
		return nil
	}
	output := make([]int, max-min)
	j := 0
	for i := min; i < max; i++ {
		output[j] = i
		j++
	}
	return output
}
