package piscine

func ConcatParams(args []string) string {
	var output string
	for i := 0; i < len(args); i++ {
		output += args[i]
		if len(args)-1 != i {
			output += "\n"
		}
	}
	return output
}
