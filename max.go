package piscine

func Max(a []int) int {
	var max int = a[0]
	for _, e := range a {
		if max < e {
			max = e
		}
	}
	return max
}
