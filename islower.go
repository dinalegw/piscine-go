package piscine

func IsLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
