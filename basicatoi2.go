package piscine

func BasicAtoi2(s string) int {
	result := 0

	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		result = result*10 + int(r-'0')
	}

	return result
}
