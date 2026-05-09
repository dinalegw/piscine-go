package piscine

func AtoiBase(s string, base string) int {
	// Validate base
	if !validBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, r := range s {
		index := indexInBase(r, base)
		if index == -1 { // invalid character
			return 0
		}
		result = result*baseLen + index
	}

	return result
}

func validBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)

	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}

	return true
}

func indexInBase(r rune, base string) int {
	for i, br := range base {
		if r == br {
			return i
		}
	}
	return -1
}
