package piscine

func AppendRange(min, max int) []int {
	var output []int
	if min == max || min > max {
		return nil
	}
	for i := min; i < max; i++ {
		output = append(output, i)
	}
	return output
}
