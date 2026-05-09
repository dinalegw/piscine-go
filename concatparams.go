package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	result := ""

	for i := 0; i < len(args); i++ {
		result += args[i]
		if i < len(args)-1 {
			result += "\n"
		}
	}

	return result
}
