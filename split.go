package piscine

func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)
	current := ""

	for i := 0; i < len(s); {
		// If the substring matches the separator
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			result = append(result, current)
			current = ""
			i += sepLen // skip the separator
		} else {
			current += string(s[i])
			i++
		}
	}

	// Add the last collected word
	result = append(result, current)

	return result
}
