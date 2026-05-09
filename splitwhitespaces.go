package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	word := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		// Check if c is a whitespace separator
		if c == ' ' || c == '\t' || c == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			// Build the current word
			word += string(c)
		}
	}

	// Add last word if exists
	if word != "" {
		result = append(result, word)
	}

	return result
}
