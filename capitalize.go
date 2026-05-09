package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, r := range runes {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			if newWord {
				// Capitalize letters at the start of a word
				if r >= 'a' && r <= 'z' {
					runes[i] = r - ('a' - 'A')
				}
				newWord = false
			} else {
				// Lowercase letters inside a word
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + ('a' - 'A')
				}
			}
		} else {
			// Non-alphanumeric character → next character starts a new word
			newWord = true
		}
	}
	return string(runes)
}
